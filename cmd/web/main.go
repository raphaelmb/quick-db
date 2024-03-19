package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/raphaelmb/quick-db/internal/database"
	"github.com/raphaelmb/quick-db/internal/sdk"
)

//go:embed template/*
var templateFS embed.FS

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /create", create)
	mux.HandleFunc("GET /list", list)
	mux.HandleFunc("POST /remove", remove)

	fmt.Println("Open on browser http://localhost:9000")
	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	dbsys := r.FormValue("dbs")
	user := r.FormValue("user")
	password := r.FormValue("password")
	port := r.FormValue("port")
	db := r.FormValue("db")
	name := r.FormValue("name")

	d := createDB(dbsys, user, password, port, db, name)

	err := sdk.Setup(d)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func remove(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	err := sdk.Remove(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "template.html")
}

func render(w http.ResponseWriter, t string) {
	tmpl, err := template.ParseFS(templateFS, fmt.Sprintf("template/%s", t))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createDB(dbsys, user, password, port, db, name string) database.DB {
	switch dbsys {
	case "postgres":
		return database.NewPostgreSQL("postgres", user, password, port, db, name, false)
	case "mysql":
		return database.NewMySQL("mysql", user, password, port, db, name, false)
	case "mongodb":
		return database.NewMongoDB("mongodb", user, password, port, db, name, false)
	}
	return nil
}
