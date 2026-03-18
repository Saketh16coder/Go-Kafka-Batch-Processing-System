package model

type Transaction struct {
	UserID int    `json:"user_id"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}
