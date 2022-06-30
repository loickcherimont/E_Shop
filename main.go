package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID       int64
	Title    string
	Img      string
	Price    float32
	Quantity int64
}

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.page.tmpl"))
}

func main() {
	setupDatabase()
	router := mux.NewRouter()
	// Serve CSS, Javascript files and images
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.HandleFunc("/", indexHandler).Methods("GET")
	log.Print("Server starting and running on 3000...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
