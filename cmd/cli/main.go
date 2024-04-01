package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/raphaelmb/quick-db/internal/database"
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
		fmt.Println(container)
	case mysql:
		mysql := database.NewMySQL("mysql", user, password, db, port, name, false)
		container, err := sdk.Setup(mysql)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		fmt.Println(container)
	case mongodb:
		mongo := database.NewMongoDB("mongo", user, password, db, port, name, false)
		container, err := sdk.Setup(mongo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(container)
	case list:
		list, err := sdk.List()
		if len(list) == 0 {
			fmt.Println("No containers found")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(list)
		return
	case remove != "":
		err := sdk.Remove(remove)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Container %s removed\n", remove)
	}
}
