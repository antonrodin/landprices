package handlers

import (
	"log"
	"net/http"

	"github.com/antonrodin/landprices/internal/models"
	"github.com/antonrodin/landprices/internal/models/mysqlite"
	"github.com/go-chi/chi/v5"
)

var App *AppRepository

type AppRepository struct {
	Transaction *mysqlite.TransactionModel
}

func NewRepo(a *AppRepository) {
	App = a
}

func (app *AppRepository) Home(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Error:   false,
		Message: "Service is alive",
	}

	app.writeJSON(w, 200, resp)
}

// Show By ID route
func (app *AppRepository) Show(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Error:   false,
		Message: "Test route",
	}

	// Get the ID from the URL
	id := chi.URLParam(r, "id")

	log.Println("ID", id)

	tr, err := app.Transaction.Get(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp.Data = tr

	app.writeJSON(w, 200, resp)
}

// Search route
func (app *AppRepository) Search(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Error:   false,
		Message: "Test route",
	}

	var all []models.Transaction

	// Get the request json data from the body
	var requestPayload struct {
		Postcode string `json:"postcode,omitempty"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	all, err = app.Transaction.Search(requestPayload.Postcode)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp.Data = all

	app.writeJSON(w, 200, resp)
}
