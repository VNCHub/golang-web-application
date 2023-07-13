package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func DataBaseConect() *sql.DB {
	conexao := "user=vinicius dbname=alura_loja password=senior host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := DataBaseConect()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := DataBaseConect()
	selection, err := db.Query("SELECT * FROM Produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	products := []Produto{}

	for selection.Next() {
		err := selection.Scan(
			&p.Id,
			&p.Nome,
			&p.Descricao,
			&p.Preco,
			&p.Quantidade,
		)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
