package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/antonrodin/landprices/internal/handlers"
	"github.com/antonrodin/landprices/internal/models/mysqlite"
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

	db, err := sql.Open("sqlite3", "./database/prices.db")
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
		Addr:    ":3333",
		Handler: app.routes(),
	}

	log.Println("Server is running on http://localhost:3333")

	return server.ListenAndServe()
}
