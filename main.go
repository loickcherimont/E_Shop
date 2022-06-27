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
	// Setup database and use it
	setupDatabase()
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	log.Print("Server starting and running on 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
