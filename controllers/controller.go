package controllers

import (
	"golang-web-application/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	product := models.SearchProducs()
	temp.ExecuteTemplate(w, "Index", product)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))
		models.NewProduct(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
