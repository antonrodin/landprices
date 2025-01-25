package mysqlite

import (
	"database/sql"

	"github.com/antonrodin/landprices/internal/models"
)

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Search(postcode string) ([]models.Transaction, error) {
	sql := `SELECT id, price, date, postcode, locality, town_city 
			FROM transactions
			WHERE postcode = ?`

	transactions := []models.Transaction{}

	rows, err := m.DB.Query(sql, postcode)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tr := models.Transaction{}
		err := rows.Scan(&tr.Id, &tr.Price, &tr.Date, &tr.Postcode, &tr.Locality, &tr.TownCity)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tr)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (m *TransactionModel) All() ([]models.Transaction, error) {
	sttm := `
		SELECT id, price, date, locality, town_city
		FROM transactions 
		LIMIT 10
		`

	rows, err := m.DB.Query(sttm)
	if err != nil {
		return nil, err
	}

	transactions := []models.Transaction{}
	for rows.Next() {
		tr := models.Transaction{}
		err := rows.Scan(&tr.Id, &tr.Price, &tr.Date, &tr.Postcode, &tr.Locality, &tr.TownCity)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tr)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
