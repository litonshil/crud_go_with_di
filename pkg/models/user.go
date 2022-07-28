package models

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Type     string `json:"type"`
	Password string `json:"password"`
}

func checkUserType(value interface{}) error {
	typ, _ := value.(string)
	if typ != "normal-user" {
		return errors.New("must be normal-user / admin")
	}
	return nil
}

func (a User) Validate() error {
	return validation.ValidateStruct(&a,
		// Name cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Type, validation.Required, validation.By(checkUserType)),
		validation.Field(&a.Password, validation.Required, validation.Length(3, 20)),
	)
}
