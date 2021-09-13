package models

import "time"

type Account struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Balance struct {
	Amount float64 `json:"balance"`
}

type Accounts []Account