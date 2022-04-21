package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store/teststore"
	"testing"
)

//TestServer_AuthenticateUser тестируем с нашим teststore. Из-за наличия testore мы не тревожим нашу бд без надобности.
//И еще тесты можно запускать параллельно с помощью go t.Run(... Оставим задачу параллельного запуска на потом.
//Хорошо читаемые табличные тесты.
func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)
	s := model.NewSession(u)
	store.Session().Create(s)

	testCases := []struct {
		name         string
		header       map[string]string
		expectedCode int
	}{
		{
			name: "authenticated",
			header: map[string]string{
				"X-Token": s.Token,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			header:       nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	server := newServer(store)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			for k, v := range tc.header {
				req.Header.Set(k, v)
			}
			server.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    "new_login",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid body",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    u.Login,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "wrong login",
			payload: map[string]string{
				"login":    "wrong_login",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "wrong password",
			payload: map[string]string{
				"login":    u.Login,
				"password": "wrong_password",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)
			s.logger.Info(rec.Body)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
