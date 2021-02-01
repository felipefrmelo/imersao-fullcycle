package model

import (
	valid "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID    string `json:"id" valid:"uuid"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"number" valid:"email"`
}

func (user *User) isValid() error {
	_, err := valid.ValidateStruct(user)

	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*User, error) {

	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	err := user.isValid()

	if err != nil {
		return nil, err
	}
	return &user, nil
}
