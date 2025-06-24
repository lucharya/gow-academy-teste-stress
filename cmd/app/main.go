package main

import (
	"log"
	"net/http"
	"os"

	"gow-academy-tst-stress/internal/db"
	"gow-academy-tst-stress/internal/programador"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" {
		log.Fatal("Uma ou mais variáveis de ambiente não foram definidas.")
	}

	database, err := db.Connect(dbUser, dbPass, dbHost, dbName)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	database.SetMaxOpenConns(100)
	database.SetMaxIdleConns(100)
	defer database.Close()

	handler := programador.NewHandler(database)

	http.HandleFunc("/programadores", handler.CadastrarProgramador)
	http.HandleFunc("/contagem-programadores", handler.ContarProgramadores)

	log.Println("Servidor iniciado na porta :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
