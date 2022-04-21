package apiserver

import (
	"encoding/json"
	"net/http"
	"some-go-api/internal/app/model"
)

//handleUsersCreate Находим в теле запроса параметры login и password. Создаем пользователя.
//В случае, если такой логин существует, то из функции создания придет ошибка.
//Возвращаем обратно модель созданного пользователя. Очищаем приватные поля пользователя перед отправкой.
func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Login    string `json: "login"`
		Password string `json: "password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Login:    req.Login,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

//handleSessionsCreate Создание сессии. Находим Логин/Пароль в теле запроса. Ищем пользователя с логином и проверяем пароль.
//Если все верно, то создаем сессию и записываем в ответ ее токен.
//Здесь же логируем результат, если пользователь с таким именем нашелся.
func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Login    string `json: "login"`
		Password string `json: "password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByLogin(req.Login)

		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errIncorrectLoginOrPassword)
			return
		}

		attemptCount, err := s.store.AuthLog().FailedAttemptsCount(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		e := &model.AuthenticationLog{
			UserID: u.ID,
			Event:  model.AuthorizeSuccess,
		}

		//логирование результата аутентификации
		defer func() {
			err := s.store.AuthLog().LogAuthenticateAttempt(e)
			if err != nil {
				s.logger.Warn("LogAuthenticateAttempt error", err)
			}
		}()

		if attemptCount >= 5 {
			e.Event = model.AuthorizeBlockedUser
			s.error(w, r, http.StatusUnauthorized, errUserBlocked)
			return
		}

		if u.ComparePassword(req.Password) == false {
			s.error(w, r, http.StatusUnauthorized, errIncorrectLoginOrPassword)
			e.Event = model.AuthorizeWrongPassword
			return
		}
		session := model.NewSession(u)
		err = s.store.Session().Create(session)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, session)
	}
}

//handleGetAuthenticateHistory Берем пользователя из контекста запроса. Он должен быть добавлен с помощью authenticateUser
//Получаем историю попыток входа этого пользователя
func (s *server) handleGetAuthenticateHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		history, err := s.store.AuthLog().GetAuthenticateHistory(r.Context().Value(ctxKeyUser).(*model.User))
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, map[string][]*model.AuthenticationLog{"data": history})
	}
}

//handleDeleteHistory Берем пользователя из контекста запроса. Он должен быть добавлен с помощью authenticateUser.
//Удаляем историю попыток входа этого пользователя
func (s *server) handleDeleteHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.store.AuthLog().DeleteAuthorizeHistory(r.Context().Value(ctxKeyUser).(*model.User))
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, map[string]string{"result": "all history deleted"})
	}
}
