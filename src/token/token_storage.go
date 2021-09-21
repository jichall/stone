package token

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/jichall/stone/src/models"
)

type TokenStorage map[int64]*models.Token

func New() TokenStorage {
	return make(TokenStorage)
}

func (ts TokenStorage) Create(id int64, expiration time.Duration, auth models.Authentication) *models.Token {

	hasher := sha1.New()
	hasher.Write([]byte(auth.CPF + auth.Secret + time.Now().String()))

	token := hex.EncodeToString(hasher.Sum(nil)[:16])

	// create a token entity and save it in the storage
	t := &models.Token{
		Token:      token,
		Expiration: expiration * time.Second,
		Issued:     time.Now(),
	}

	ts[id] = t

	return t
}

func (ts TokenStorage) IsValid(token string) bool {

	var temp *models.Token

	for _, t := range ts {
		if t.Token == token {
			temp = t
			break
		}
	}

	if temp == nil {
		return false
	}

	if temp.Issued.Add(temp.Expiration * time.Second).Before(time.Now()) {
		return false
	}

	return true
}
