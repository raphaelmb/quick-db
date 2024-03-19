package sdk

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func List() ([]Container, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	filter := filters.NewArgs()
	filter.Add("label", "quickdb=generated")

	containers, err := cli.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filter,
	})
	if err != nil {
		return []Container{}, fmt.Errorf("error listing containers: %w", err)
	}

	var result []Container
	for _, container := range containers {
		result = append(result, toContainer(container))
	}

	fmt.Println(result)
	return result, nil
}

type Container struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Port  uint16 `json:"port"`
}

func toContainer(container types.Container) Container {
	return Container{
		ID:    container.ID,
		Name:  container.Names[0],
		Image: container.Image,
		Port:  container.Ports[0].PublicPort,
	}
}
