package model

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

//TestUser Тестовый пользователь для того чтобы не дублировать код в тестах.
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

