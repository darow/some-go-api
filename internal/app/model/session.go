package model

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

//Тут несколько вариантов времени жизни сессии. Подольше и покороче для тестирования.
const (
	sessionLiveTimeShort = time.Second * time.Duration(20)
	sessionLiveTime = time.Minute * time.Duration(30)
	sessionLiveTimeLong = time.Hour * time.Duration(1000)
)

type Session struct {
	ID             int       `json:"-"`
	UserID         int       `json:"user_id"`
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"exp"`
}

//CreateToken вспомогательная функция создания токена для сессии. У каждой сессии должен быть уникальный ExpirationTime.UnixNano().
func (s *Session) CreateToken() {
	b := md5.Sum([]byte(strconv.FormatInt(s.ExpirationTime.UnixNano(), 10)))
	token := hex.EncodeToString(b[:])
	s.Token = token[:]
}

//NewSession Создание сессии и вызов создания токена для нее. Здесь нужно менять sessionLiveTime
func NewSession(u *User) *Session {
	s := &Session{
		UserID:         u.ID,
		ExpirationTime: time.Now().Local().Add(sessionLiveTimeLong),
	}
	s.CreateToken()
	return s
}
