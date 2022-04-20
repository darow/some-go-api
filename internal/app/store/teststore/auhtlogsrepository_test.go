package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store/teststore"
	"testing"
)

func TestAuthLogRepository_LogAuthenticateAttempt(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	log := &model.AuthenticationLog{
		Event: model.AuthorizeSuccess,
		UserID: testUser.ID,
	}

	err := s.AuthLog().LogAuthenticateAttempt(log)
	assert.NoError(t, err)
}

func TestAuthLogRepository_FailedAttemptsCount(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	log := &model.AuthenticationLog{
		Event: model.AuthorizeWrongPassword,
		UserID: testUser.ID,
	}
	s.AuthLog().LogAuthenticateAttempt(log)

	n, err := s.AuthLog().FailedAttemptsCount(testUser)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	log.Event = model.AuthorizeSuccess
	s.AuthLog().LogAuthenticateAttempt(log)
	n, err = s.AuthLog().FailedAttemptsCount(testUser)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	log.Event = model.AuthorizeWrongPassword
	s.AuthLog().LogAuthenticateAttempt(log)
	n, err = s.AuthLog().FailedAttemptsCount(testUser)
	assert.NoError(t, err)
	assert.Equal(t, 2, n)
}


func TestAuthLogRepository_GetAuthenticateHistory(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	log := &model.AuthenticationLog{
		Event: model.AuthorizeSuccess,
		UserID: testUser.ID,
	}
	s.AuthLog().LogAuthenticateAttempt(log)

	logs, err := s.AuthLog().GetAuthenticateHistory(testUser)
	assert.NoError(t, err)
	assert.Equal(t, []*model.AuthenticationLog{log}, logs)
}

func TestAuthLogRepository_DeleteAuthorizeHistory(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	log := &model.AuthenticationLog{
		Event: model.AuthorizeSuccess,
		UserID: testUser.ID,
	}
	s.AuthLog().LogAuthenticateAttempt(log)

	err := s.AuthLog().DeleteAuthorizeHistory(testUser)
	assert.NoError(t, err)
	logs, _ := s.AuthLog().GetAuthenticateHistory(testUser)
	assert.Equal(t, []*model.AuthenticationLog{}, logs)
}