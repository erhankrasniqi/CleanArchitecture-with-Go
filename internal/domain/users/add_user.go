package users

import "fmt"

type AddUser struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (au *AddUser) Validate() error {
	if au.UserName == "" {
		return fmt.Errorf("username is required")
	}
	if au.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
