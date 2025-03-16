package usecases

import (
	"firstProject/internal/domain/users"
	db "firstProject/internal/infrastructure/repository"
)

type UserUsecase struct {
	UserRepo *db.UserRepository
}

func NewUserUsecase(userRepo *db.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (u *UserUsecase) GetAllUsers() ([]users.User, error) {
	users, err := u.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
