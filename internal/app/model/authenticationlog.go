package model

import (
	"encoding/json"
	"time"
)

const (
	AuthorizeSuccess = iota
	AuthorizeWrongPassword
	AuthorizeBlockedUser
)

//AuthenticationLog модель для лога результата аутентификации. Используется как для внесения в бд, так и для чтения истории логов из бд
type AuthenticationLog struct {
	Timestamp time.Time `json:"timestamp"`
	UserID    int       `json:"-"`
	Event     uint8     `json:"event"`
}

//MarshalJSON переопределяем функцию формирования json из структуры. Меняем формат времени и формируем понятное название
//события аутентификации
func (l *AuthenticationLog) MarshalJSON() ([]byte, error) {
	events := map[uint8]string{
		0: "AuthorizeSuccess",
		1: "AuthorizeWrongPassword",
		2: "AuthorizeBlockedUser",
	}
	type Alias AuthenticationLog
	return json.Marshal(&struct {
		*Alias
		Timestamp string `json:"timestamp"`
		Event     string `json:"event"`
	}{
		Alias:     (*Alias)(l),
		Timestamp: l.Timestamp.Format("2006/01/02 15:04:05"),
		Event:     events[l.Event],
	})
}
