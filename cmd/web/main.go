package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//go:embed template/*
var templateFS embed.FS

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /create", create)
	mux.HandleFunc("GET /list", list)
	mux.HandleFunc("POST /remove", remove)

	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func list(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func remove(w http.ResponseWriter, r *http.Request) {
	// TODO
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
