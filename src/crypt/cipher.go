package crypt

import "golang.org/x/crypto/bcrypt"

func Encrypt(secret string) string {

	b, err := bcrypt.GenerateFromPassword([]byte(secret), 14)

	if err != nil {
		panic(err)
	}

	return string(b)
}
