package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/raphaelmb/quick-db/internal/database"
	"github.com/raphaelmb/quick-db/internal/dto"
	"github.com/raphaelmb/quick-db/internal/sdk"
)

func main() {
	var (
		postgres bool
		mysql    bool
		mongodb  bool
		list     bool
		remove   string

		user     string
		password string
		db       string
		port     string
		name     string
	)

	flag.BoolVar(&postgres, "postgres", false, "create postgres database")
	flag.BoolVar(&mysql, "mysql", false, "create mysql database")
	flag.BoolVar(&mongodb, "mongodb", false, "create mongodb")

	flag.BoolVar(&list, "list", false, "list containers")
	flag.StringVar(&remove, "remove", "", "remove container")

	flag.StringVar(&user, "user", "", "database user")
	flag.StringVar(&password, "password", "", "database password")
	flag.StringVar(&db, "db", "", "default database")
	flag.StringVar(&port, "port", "", "host port")
	flag.StringVar(&name, "name", "", "container name")

	flag.Parse()

	switch {
	case postgres:
		pg := database.NewPostgreSQL("postgres", user, password, db, port, name, false)
		container, err := sdk.Setup(pg)
		if err != nil {
			log.Fatal(err)
		}
		PrintContainer(container)
	case mysql:
		mysql := database.NewMySQL("mysql", user, password, db, port, name, false)
		container, err := sdk.Setup(mysql)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		PrintContainer(container)
	case mongodb:
		mongo := database.NewMongoDB("mongo", user, password, db, port, name, false)
		container, err := sdk.Setup(mongo)
		if err != nil {
			log.Fatal(err)
		}
		PrintContainer(container)
	case list:
		list, err := sdk.List()
		if len(list) == 0 {
			fmt.Println("No containers found")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		ListContainers(list)
	case remove != "":
		err := sdk.Remove(remove)
		if err != nil {
			fmt.Printf("No container with id %s found\n", remove)
			os.Exit(0)
		}
		fmt.Printf("Container %s removed\n", remove)
	}
}

func PrintContainer(container dto.ContainerCreate) {
	fmt.Printf("ID: %s\nName: %s\nPort: %s\nUser: %s\nPassword: %s\nDatabase: %s\nDSN: %s\n",
		container.ID, container.Name, container.Port, container.User, container.Password, container.Database, container.DSN)
}

func ListContainers(containers []dto.ContainerList) {
	for _, v := range containers {
		fmt.Printf("ID: %s\nName: %s\nImage: %s\nPort: %d\n\n", v.ID, v.Name, v.Image, v.Port)
	}
}
