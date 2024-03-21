package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Heading string
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		data := PageData{
			Title:   "Hello, Kubernetes!",
			Heading: "Welcome to the Go Web Server",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
