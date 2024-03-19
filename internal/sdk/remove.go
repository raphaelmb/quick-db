package sdk

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Remove(id string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	defer cli.Close()

	con, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return fmt.Errorf("error inspecting container: %w", err)
	}

	if err := cli.ContainerRemove(ctx, con.ID, container.RemoveOptions{Force: true}); err != nil {
		return fmt.Errorf("error removing container: %w", err)
	}

	return nil
}
