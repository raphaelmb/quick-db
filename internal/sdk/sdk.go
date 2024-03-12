package sdk

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func cli() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func pull(ctx context.Context, cli *client.Client, image string) (io.ReadCloser, error) {
	url := fmt.Sprintf("docker.io/library/%s", image)
	reader, err := cli.ImagePull(ctx, url, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func create(ctx context.Context, cli *client.Client, image string) (container.CreateResponse, error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, nil, nil, nil, "")
	if err != nil {
		return container.CreateResponse{}, err
	}
	return resp, nil
}

func start(ctx context.Context, cli *client.Client, resp container.CreateResponse) error {
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return err
	}
	return nil
}

func Setup(image string) {
	ctx := context.Background()
	cli, err := cli()
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err := pull(ctx, cli, image)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	io.Copy(io.Discard, reader)

	resp, err := create(ctx, cli, image)
	if err != nil {
		panic(err)
	}

	if err := start(ctx, cli, resp); err != nil {
		panic(err)
	}
}
