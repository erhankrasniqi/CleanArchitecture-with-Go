package users

import "github.com/go-playground/validator/v10"

type UpdateUser struct {
	UserName string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
}

func (u *UpdateUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
