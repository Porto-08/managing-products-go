package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectWithDatabase() *sql.DB {
	conection := "user=postgres dbname=alura_loja password=samuel host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}

	print("Conectado ao banco com sucesso!")
	return db
}
