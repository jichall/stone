package models

type Authentication struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Token struct {
	Token string `json:"token"`
}