package usecases

import (
	"firstProject/internal/domain/users"
	db "firstProject/internal/infrastructure/repository"
)

type UserByIDUsecase struct {
	UserRepo *db.UserRepository
}

func NewUserByIDUsecase(userRepo *db.UserRepository) *UserByIDUsecase {
	return &UserByIDUsecase{UserRepo: userRepo}
}

func (u *UserByIDUsecase) GetUserByID(id int) (*users.User, error) {
	return u.UserRepo.GetUserByID(id)
}
