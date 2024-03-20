package sdk

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/raphaelmb/quick-db/internal/dto"
)

func List() ([]dto.ContainerList, error) {
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
		return []dto.ContainerList{}, fmt.Errorf("error listing containers: %w", err)
	}

	var result []dto.ContainerList
	for _, container := range containers {
		result = append(result, dto.ToContainerListDTO(container))
	}

	return result, nil
}
