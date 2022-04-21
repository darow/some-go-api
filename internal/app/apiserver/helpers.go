package apiserver

import (
	"encoding/json"
	"net/http"
)

//error Запись в ResponseWriter ошибки
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

//respond Записываем в заголовок ответа код, соответствующий результату. И создаем json для записи в ResponseWriter
//r *http.Request получаем на всякий случай. Вдруг еще пригодится
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
