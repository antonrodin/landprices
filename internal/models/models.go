package models

type Transaction struct {
	Id       string `json:"id"`
	Price    int    `json:"price"`
	Date     string `json:"date"`
	Postcode string `json:"postcode"`
	TownCity string `json:"townCity"`
	Locality string `json:"locality"`
}
