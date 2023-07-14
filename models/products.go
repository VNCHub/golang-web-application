package models

import "golang-web-application/db"

type Products struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SearchProducs() []Products {
	db := db.DataBaseConect()
	selection, err := db.Query("SELECT * FROM Produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Products{}
	products := []Products{}

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
	return products
}
