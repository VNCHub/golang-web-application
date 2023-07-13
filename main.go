package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func DataBaseConect() *sql.DB {
	conexao := "username=vinicius dbname=alura_loja password=senior host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Tenis", "Confortavel", 89, 120},
		{"Computador", "Bem rápido", 3999, 10},
		{"Celular", "Resistênte", 299, 50},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
