package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-,omitempty"`
}

//BeforeCreate Метод должен вызываться перед записью в БД. Тут можно делать не только шифрование пароля.
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

//Sanitize Чистим приватные поля. Например записью найденного пользователя в ответ.
func (u *User) Sanitize() {
	u.Password = ""
	u.EncryptedPassword = ""
}

//ComparePassword метод сравнения хеша и пароля. Если надо будет менять способ шифрования, то будем менять здесь.
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

//encryptString Шифрование строки, содержащей пароль
func encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
