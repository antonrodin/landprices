package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antonrodin/landprices/internal/handlers"
	"github.com/antonrodin/landprices/internal/models/mysqlite"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Title string
}

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Config
	port := os.Getenv("PORT")
	dbFile := os.Getenv("DB_FILE")

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("Connected to SQLite successfully.")

	app := Config{
		Title: "Land Prices API",
	}

	appRepo := handlers.AppRepository{
		Transaction: &mysqlite.TransactionModel{
			DB: db,
		},
	}

	handlers.NewRepo(&appRepo)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	log.Printf("Listening on port %s", port)

	return server.ListenAndServe()
}
