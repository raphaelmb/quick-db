package main

import (
	"embed"
	"encoding/json"
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
	mux.HandleFunc("DELETE /remove/{id}", remove)

	fmt.Println("Open on browser http://localhost:9000")
	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}

type Input struct {
	Dbs      string `json:"dbs"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

func create(w http.ResponseWriter, r *http.Request) {
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := createDB(input.Dbs, input.User, input.Password, input.Db, input.Port, input.Name)

	c, err := sdk.Setup(d)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(m)
}

func list(w http.ResponseWriter, r *http.Request) {
	c, err := sdk.List()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}

func remove(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

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
