package controllers

import (
	"golang-web-application/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchProducs()
	temp.ExecuteTemplate(w, "Index", products)
}
