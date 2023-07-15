package models

import "golang-web-application/db"

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SearchProducs() []Product {
	db := db.DataBaseConect()
	selection, err := db.Query("SELECT * FROM Produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

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

func NewProduct(nome string, descricao string, preco float64, quantidade int) {

	db := db.DataBaseConect()
	selection, err := db.Prepare(
		"INSERT INTO Produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)",
	)
	if err != nil {
		panic(err.Error())
	}
	selection.Exec(nome, descricao, preco, quantidade)
}

func DeleteProduct(productId int) {
	db := db.DataBaseConect()
	selection, err := db.Prepare(
		"DELETE FROM Produtos WHERE Id = $1",
	)
	if err != nil {
		panic(err.Error())
	}
	selection.Exec(productId)
	defer db.Close()
}

func EditProduct(ID string) Product {
	db := db.DataBaseConect()
	query, err := db.Query("SELECT * FROM Produtos WHERE Id =$1", ID)
	if err != nil {
		panic(err.Error())
	}
	p := Product{}

	for query.Next() {
		query.Scan(
			&p.Id,
			&p.Nome,
			&p.Descricao,
			&p.Preco,
			&p.Quantidade,
		)
	}

	defer db.Close()
	return p
}
