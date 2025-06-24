package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(user, password, host, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS programadores (
		id UUID PRIMARY KEY,
		apelido VARCHAR(32) UNIQUE NOT NULL,
		nome VARCHAR(100) NOT NULL,
		nascimento VARCHAR(10) NOT NULL,
		stack TEXT
	);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Printf("Erro criando tabela programadores: %v", err)
		return err
	}

	return nil
}
