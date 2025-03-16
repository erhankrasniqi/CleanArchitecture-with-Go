package users

import "github.com/go-playground/validator/v10"

type DeleteUser struct {
	ID int `json:"id" validate:"required,gt=0"`
}

func (d *DeleteUser) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
