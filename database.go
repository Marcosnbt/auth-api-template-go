// Arquivo criado em maio
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver para PostgreSQL, pode ser outro se quiser
)

var db *sql.DB

func InitDB() {
	connStr := "user=usuario dbname=nomedb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Conexão com banco de dados estabelecida.")
}
