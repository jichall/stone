package models

import "time"

type Authentication struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Token struct {
	Token string `json:"token"`
	Issued time.Time `json:"issued"`
	Expiration time.Duration `json:"expiration"`
}