package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"some-go-api/internal/app/store/sqlstore"
	"testing"
)

//TestUserRepository_Create Проверяем, что при создании пользователя не возвращается ошибок
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


