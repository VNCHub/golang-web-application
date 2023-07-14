package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DataBaseConect() *sql.DB {
	conexao := "user=vinicius dbname=alura_loja password=senior host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
