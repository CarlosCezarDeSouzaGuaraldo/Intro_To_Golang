package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name string
	Occupation string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		u := user{"Carlos", "Developer"}

		templates.ExecuteTemplate(w, "index.html", u)
	})

	fmt.Println("Listening on PORT: 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
