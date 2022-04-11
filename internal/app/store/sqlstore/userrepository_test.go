package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"some-go-api/internal/app/store/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("users")
	s := sqlstore.New(db)
	assert.NoError(t, s.User().Create(model.TestUser(t)))
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("users")
	s := sqlstore.New(db)

	testUser := model.TestUser(t)
	_, err := s.User().Find(testUser.ID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.User().Create(testUser)

	foundUser, err := s.User().Find(testUser.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, testUser.Login == foundUser.Login, true)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("users")
	s := sqlstore.New(db)

	testUser := model.TestUser(t)
	_, err := s.User().FindByLogin(testUser.Login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.User().Create(testUser)

	foundUser, err := s.User().FindByLogin(testUser.Login)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, testUser.Login == foundUser.Login, true)
}

//func TestUserRepository_CheckPass(t *testing.T) {
//	db, teardown := sqlstore.TestDB(t, psqlInfo)
//	defer teardown("users")
//	s := sqlstore.New(db)
//
//	testUser := model.TestUser(t)
//	s.User().Create(testUser)
//
//	u, err := s.User().FindByLogin(testUser.Login)
//	match, err := s.User().CheckPass(u, testUser.Password)
//
//	assert.NoError(t, err)
//	assert.True(t, match)
//}

func TestUserRepository_FindByLoginPass(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("users")
	s := sqlstore.New(db)

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	foundedUser, err := s.User().FindByLoginPass(testUser.Login, testUser.Password)
	assert.NoError(t, err)
	assert.NotNil(t, foundedUser)
	assert.Equal(t, testUser.Login == foundedUser.Login, true)
}

