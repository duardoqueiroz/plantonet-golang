package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func getConnection() (*sqlx.DB, error) {
	url := "postgres://postgres:12345678@localhost:5432/plantonetgo?sslmode=disable"
	conn, err := sqlx.Open("postgres", url)

	if err != nil {
		log.Println("Houve um erro ao conectar com o banco de dados")
		return nil, err
	}

	return conn, nil
}

func closeConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Println(err)
	}
}
