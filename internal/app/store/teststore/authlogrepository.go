package teststore

import "some-go-api/internal/app/model"

type AuthLogRepository struct {
	store *Store
	authLogs map[int]*model.AuthorizationLog
}

func (r *AuthLogRepository) LogAuthenticateAttempt(event *model.AuthorizationLog) error {
	id := len(r.authLogs)
	r.authLogs[id] = event
	return nil
}

func (r *AuthLogRepository) FailedAttemptsCount(u *model.User) (count int, err error) {
	for _, v := range r.authLogs {
		if v.UserID == u.ID {
			if v.Event == model.AuthorizeWrongPassword {
				count++
			}
		}
	}
	return  count, nil
}

func (r *AuthLogRepository) GetAuthorizeHistory(u *model.User) (logs []*model.AuthorizationLog, err error) {
	//TODO
	return logs, nil
}

func (r *AuthLogRepository) DeleteAuthorizeHistory(u *model.User) error {
	//TODO
	return nil
}