package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Session struct {
	ID             int       `json: "id"`
	UserID         int       `json: "user_id"`
	Token          string    `json: "token"`
	ExpirationTime time.Time `json: "exp"`
}

func (s *Session) CreateToken() {
	logrus.Info("CreateToken", s.ExpirationTime.UnixNano())
	b := md5.Sum([]byte(strconv.FormatInt(s.ExpirationTime.UnixNano(), 10)))
	token := hex.EncodeToString(b[:])
	logrus.Info("CreateToken", token)
	s.Token = token[:]
}
