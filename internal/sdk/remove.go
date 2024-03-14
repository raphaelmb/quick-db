package sdk

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Remove(id string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	con, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerRemove(ctx, con.ID, container.RemoveOptions{Force: true}); err != nil {
		panic(err)
	}
}
