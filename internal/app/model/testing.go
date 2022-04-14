package model

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUser(t *testing.T) *User {
	pass := "Password1"
	enc, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	return &User{
		ID:                1,
		Login:             "test_login",
		Password:          pass,
		EncryptedPassword: string(enc),
	}
}

