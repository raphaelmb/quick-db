package sdk

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func List() {
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
		panic(err)
	}

	for _, container := range containers {
		fmt.Println("ID:", container.ID[:12])
		fmt.Println("Name:", container.Names[0])
		fmt.Println("Image:", container.Image)
		fmt.Println("Port:", container.Ports[0].PublicPort)
	}
}
