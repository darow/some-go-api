package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store/sqlstore"
	"testing"
)

func TestUserRepository_LogAuthenticateAttempt(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("users", "authorization_events")
	s := sqlstore.New(db)

	testUser := model.TestUser(t)
	s.User().Create(testUser)
	e := &model.AuthorizationLog{
		UserID: testUser.ID,
		Event: model.AuthorizeSuccess,
	}

	err := s.AuthLog().LogAuthenticateAttempt(e)
	assert.NoError(t, err)
}

func TestUserRepository_FailedAttemptsCount(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("users", "authorization_events")
	s := sqlstore.New(db)

	testUser := model.TestUser(t)
	s.User().Create(testUser)
	e := &model.AuthorizationLog{
		UserID: testUser.ID,
		Event: model.AuthorizeSuccess,
	}

	err := s.AuthLog().LogAuthenticateAttempt(e)
	assert.NoError(t, err)
}
