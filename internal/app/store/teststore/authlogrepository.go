package teststore

import (
	"some-go-api/internal/app/model"
)

type AuthLogRepository struct {
	store *Store
	authLogs []*model.AuthenticationLog
}

func (r *AuthLogRepository) LogAuthenticateAttempt(event *model.AuthenticationLog) error {
	newLog := &model.AuthenticationLog{
		Timestamp: event.Timestamp,
		UserID: event.UserID,
		Event: event.Event,
	}
	r.authLogs = append(r.authLogs, newLog)
	return nil
}

func (r *AuthLogRepository) FailedAttemptsCount(u *model.User) (count int, err error) {
	for _, l := range r.authLogs {
		if l.UserID == u.ID {
			if l.Event == model.AuthorizeWrongPassword {
				count++
			}
		}
	}
	return  count, nil
}

func (r *AuthLogRepository) GetAuthenticateHistory(u *model.User) ([]*model.AuthenticationLog, error) {
	logs := []*model.AuthenticationLog{}
	for _, l := range r.authLogs {
		if l.UserID == u.ID {
			logs = append(logs, l)
		}
	}
	return logs, nil
}

func (r *AuthLogRepository) DeleteAuthorizeHistory(u *model.User) error {
	r.authLogs = []*model.AuthenticationLog{}
	return nil
}