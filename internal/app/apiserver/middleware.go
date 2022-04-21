package apiserver

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//logRequest Выводим в консоль информацию о запросе, прошедшем через этот middleware.
func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

//setRequestID Присваиваем каждому запросу ID. Передаем его пользователю. Теперь мы можем находить каждый запрос по ID.
//Удобно, но в этом проекте не обязательно.
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

//authenticateUser Берем заголовок запроса с именем X-Token. Ищем токен в таблице сессий в БД. Если находим, то передаем
//дальше пользователя, к которому эта сессия относится.
func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Token")
		session, err := s.store.Session().Find(token)

		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errTokenNotFound)
			return
		}

		if session.ExpirationTime.Before(time.Now()) {
			s.error(w, r, http.StatusUnauthorized, errTokenExpired)
			return
		}

		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		u, err := s.store.User().Find(session.UserID)

		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		u.Sanitize()
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}