package usecases

import (
	"firstProject/internal/domain/users"
)

func (u *UserUsecase) AddUser(newUser users.AddUser) (users.User, error) {
	user := users.User{
		UserName: newUser.UserName,
		Email:    newUser.Email,
	}
	return u.UserRepo.AddUser(user)
}
