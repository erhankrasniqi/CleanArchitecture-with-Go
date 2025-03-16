package usecases

import (
	"firstProject/internal/domain/users"
	db "firstProject/internal/infrastructure/repository"
)

type UserUpdateUsecase struct {
	UserRepo *db.UserRepository
}

func NewUserUpdateUsecase(userRepo *db.UserRepository) *UserUpdateUsecase {
	return &UserUpdateUsecase{UserRepo: userRepo}
}

func (u *UserUpdateUsecase) UpdateUser(id int, updatedUser users.User) error {
	return u.UserRepo.UpdateUser(id, updatedUser)
}
