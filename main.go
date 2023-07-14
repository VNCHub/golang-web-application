package main

import (
	"golang-web-application/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchProducs()
	temp.ExecuteTemplate(w, "Index", products)
}
