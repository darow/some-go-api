package model

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

const (
	sessionLiveTimeShort = time.Second * time.Duration(20)
	sessionLiveTime = time.Minute * time.Duration(30)
	sessionLiveLong = time.Hour * time.Duration(1000)
)

type Session struct {
	ID             int       `json: "id"`
	UserID         int       `json: "user_id"`
	Token          string    `json: "token"`
	ExpirationTime time.Time `json: "exp"`
}

func (s *Session) CreateToken() {
	b := md5.Sum([]byte(strconv.FormatInt(s.ExpirationTime.UnixNano(), 10)))
	token := hex.EncodeToString(b[:])
	s.Token = token[:]
}

func NewSession(u *User) *Session {
	s := &Session{
		UserID:         u.ID,
		ExpirationTime: time.Now().Local().Add(sessionLiveLong),
	}
	s.CreateToken()
	return s
}
