package mysqlite

import (
	"database/sql"

	"github.com/antonrodin/landprices/internal/models"
)

type TransactionModel struct {
	DB *sql.DB
}

// Get by ID method
func (m *TransactionModel) Get(id string) (models.Transaction, error) {
	sql := `SELECT id, price, date, old_or_new, postcode, locality, town_city, county, street, primary_address, secondary_address
			FROM transactions
			WHERE id = ?`

	tr := models.Transaction{}

	err := m.DB.QueryRow(sql, id).
		Scan(&tr.Id, &tr.Price, &tr.Date, &tr.OldOrNew, &tr.Postcode, &tr.Locality, &tr.TownCity, &tr.County, &tr.Street, &tr.PrimaryAddress, &tr.SecondaryAddress)
	if err != nil {
		return tr, err
	}

	return tr, nil
}

func (m *TransactionModel) Search(postcode string) ([]models.Transaction, error) {
	sql := `SELECT id, price, date, old_or_new, postcode, locality, town_city, county, street, primary_address, secondary_address
			FROM transactions
			WHERE postcode = ?
			LIMIT 200`

	transactions := []models.Transaction{}

	rows, err := m.DB.Query(sql, postcode)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tr := models.Transaction{}
		err := rows.Scan(&tr.Id, &tr.Price, &tr.Date, &tr.OldOrNew, &tr.Postcode, &tr.Locality, &tr.TownCity, &tr.County, &tr.Street, &tr.PrimaryAddress, &tr.SecondaryAddress)
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
		SELECT id, price, date, old_or_new, postcode, locality, town_city, county, street, primary_address, secondary_address
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
		err := rows.Scan(&tr.Id, &tr.Price, &tr.Date, &tr.OldOrNew, &tr.Postcode, &tr.Locality, &tr.TownCity, &tr.County, &tr.Street, &tr.PrimaryAddress, &tr.SecondaryAddress)
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
