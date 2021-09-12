package models

import "time"

type Transaction struct {
	ID                 int64     `json:"id"`
	AccountOrigin      string    `json:"account_origin_id"`
	AccountDestination string    `json:"account_destination_id"`
	Amount             float64   `json:"amount"`
	CreatedAt          time.Time `json:"created_at"`
}
