package handlers

import (
	"net/http"

	"github.com/antonrodin/landprices/internal/models"
	"github.com/antonrodin/landprices/internal/models/mysqlite"
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

func (app *AppRepository) Test(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Error:   false,
		Message: "Test route",
	}

	all, err := app.Transaction.All()
	if err != nil {
		resp.Error = true
		resp.Message = err.Error()
		app.writeJSON(w, 500, resp)
		return
	}

	resp.Data = all

	app.writeJSON(w, 200, resp)
}

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
