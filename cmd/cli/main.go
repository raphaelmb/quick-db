package main

import (
	"fmt"
	"os"

	"github.com/raphaelmb/quick-db/internal/database"
	"github.com/raphaelmb/quick-db/internal/sdk"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	switch arg1 {
	case "create":
		fmt.Println("Creating a database via Docker")
	case "list":
		sdk.List()
		return
	case "remove":
		// sdk.Remove()
	default:
		fmt.Println("Error: Expected command 'create', 'list' or 'remove'")
		return
	}

	switch arg2 {
	case "postgres":
		fmt.Println("PostgreSQL chosen")
		pg := database.NewPostgreSQL("postgres", "", "", "", "", "", false)
		sdk.Setup(pg)
	case "mysql":
		fmt.Println("MySQL chosen")
		mysql := database.NewMySQL("mysql", "", "", "", "", "", false)
		sdk.Setup(mysql)
	case "mongo":
		fmt.Println("MongoDB chosen")
		mongo := database.NewMongoDB("mongo", "", "", "", "", "", false)
		sdk.Setup(mongo)
	default:
		fmt.Println("Error: expected database 'postgres', 'mysql' or 'mongo'")
		return
	}
}
