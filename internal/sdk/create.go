package sdk

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/raphaelmb/quick-db/internal/database"
	"github.com/raphaelmb/quick-db/internal/dto"
)

func cli() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	return cli, nil
}

func pull(ctx context.Context, cli *client.Client, image string) (io.ReadCloser, error) {
	url := fmt.Sprintf("docker.io/library/%s", image)
	reader, err := cli.ImagePull(ctx, url, types.ImagePullOptions{})
	if err != nil {
		return nil, fmt.Errorf("error pulling image: %w", err)
	}
	return reader, nil
}

func create(ctx context.Context, cli *client.Client, image string, env []string, containerPort, hostPort, name string, vol bool, path string) (container.CreateResponse, error) {
	containerCfg := &container.Config{
		Image:  image,
		Env:    env,
		Tty:    true,
		Labels: map[string]string{"quickdb": "generated"},
	}

	hostBiding := nat.PortBinding{
		HostIP:   "127.0.0.1",
		HostPort: hostPort,
	}

	containerBiding := nat.PortMap{nat.Port(containerPort + "/tcp"): []nat.PortBinding{hostBiding}}

	var hostCfg *container.HostConfig

	if vol {
		v, err := cli.VolumeCreate(ctx, volume.CreateOptions{
			Name: "quickdb-123", // TODO
		})
		if err != nil {
			return container.CreateResponse{}, err
		}
		hostCfg = &container.HostConfig{PortBindings: containerBiding, Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: v.Name,
				Target: path,
			},
		}}
	} else {
		hostCfg = &container.HostConfig{PortBindings: containerBiding, Mounts: []mount.Mount{}}
	}

	resp, err := cli.ContainerCreate(ctx, containerCfg, hostCfg, nil, nil, name)
	if err != nil {
		return container.CreateResponse{}, fmt.Errorf("failed to create container: %w", err)
	}
	return resp, nil
}

func start(ctx context.Context, cli *client.Client, resp container.CreateResponse) error {
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("error starting container: %w", err)
	}
	return nil
}

func Setup(db database.DB) (dto.ContainerCreate, error) {
	ctx := context.Background()
	cli, err := cli()
	if err != nil {
		return dto.ContainerCreate{}, err
	}
	defer cli.Close()

	reader, err := pull(ctx, cli, db.GetImage())
	if err != nil {
		return dto.ContainerCreate{}, nil
	}
	defer reader.Close()

	io.Copy(io.Discard, reader)

	resp, err := create(ctx, cli, db.GetImage(), db.EnvVars(), db.GetContainerPort(), db.GetHostPort(), db.GetContainerName(), db.GetCreateVolume(), db.GetDataPath())
	if err != nil {
		return dto.ContainerCreate{}, err
	}

	if err := start(ctx, cli, resp); err != nil {
		return dto.ContainerCreate{}, err
	}

	name, err := getContainerName(ctx, cli)
	if err != nil {
		return dto.ContainerCreate{}, err
	}

	return dto.ContainerCreate{
		ID:       resp.ID,
		Name:     name,
		Port:     db.GetHostPort(),
		User:     db.GetUser(),
		Password: db.GetPassword(),
		Database: db.GetDB(),
		DSN:      db.Dsn(),
	}, nil
}

func getContainerName(ctx context.Context, cli *client.Client) (string, error) {
	container, err := cli.ContainerList(ctx, container.ListOptions{Latest: true})
	if err != nil {
		return "", err
	}
	return container[0].Names[0], nil
}
