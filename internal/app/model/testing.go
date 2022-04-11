package model

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
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

func TestSession(t *testing.T) *Session {
	futureTime := time.Now().Local().Add(time.Minute * time.Duration(5))
	randomValue := string(futureTime.UnixMilli())
	token, _ := bcrypt.GenerateFromPassword([]byte(randomValue), bcrypt.MinCost)
	return &Session{
		ID:             1,
		UserID:         0,
		Token:          string(token),
		ExpirationTime: futureTime,
	}
}
