package dto

import (
	"github.com/docker/docker/api/types"
)

type ContainerList struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Port  uint16 `json:"port"`
}

func ToContainerListDTO(container types.Container) ContainerList {
	return ContainerList{
		ID:    container.ID[:7],
		Name:  container.Names[0],
		Image: container.Image,
		Port:  container.Ports[0].PublicPort,
	}
}

type ContainerCreate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	DSN      string `json:"dsn"`
}

func ToContainerCreateDTO(id, name, port, user, password, database, dsn string) ContainerCreate {
	return ContainerCreate{
		ID:       id[:7],
		Name:     name,
		User:     user,
		Password: password,
		Port:     port,
		Database: database,
		DSN:      dsn,
	}
}
