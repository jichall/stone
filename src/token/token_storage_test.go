package token

import (
	"testing"
	"time"

	"github.com/jichall/stone/src/models"
)

type Tests struct {
	ID         int64
	Auth       models.Authentication
	// this field represents in how many seconds the token will expire in a text
	Expiration time.Duration
}

var (
	tests = []Tests{
		{
			ID:         1,
			Auth:       models.Authentication{CPF: "12345678910", Secret: "123123"},
			Expiration: time.Second * 2,
		},
		{
			ID:         2,
			Auth:       models.Authentication{CPF: "12635678910", Secret: "321321"},
			Expiration: time.Second * 4,
		},
		{
			ID:         3,
			Auth:       models.Authentication{CPF: "12635672110", Secret: "321453"},
			Expiration: time.Second * 1,
		},
	}
)

func TestStorage(t *testing.T) {

	storage := New()

	for _, test := range tests {

		t := storage.Create(test.ID, test.Expiration, test.Auth)

		storage.IsValid(t.Token)
	}
}
