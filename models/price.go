package models

type Price struct {
	ID     uint    `json:"id"`
	Item   string  `json:"item"`
	Store  string  `json:"store"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}
