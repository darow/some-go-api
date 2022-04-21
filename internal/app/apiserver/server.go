package apiserver

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"some-go-api/internal/app/store"
)

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectLoginOrPassword = errors.New("incorrect login or password")
	errNotAuthenticated         = errors.New("not authenticated")
	errUserBlocked              = errors.New("user blocked")
	errTokenExpired             = errors.New("token has expired")
	errTokenNotFound            = errors.New("token not found")
)

type ctxKey uint8

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

//newServer Создаем экземпляр сервера. Лаконично и здорово!
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	return s
}

//ServeHTTP более краткий вызов функции обслуживание запроса.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/sessions", s.handleGetAuthenticateHistory()).Methods("GET")
	private.HandleFunc("/sessions", s.handleDeleteHistory()).Methods("DELETE")
}
