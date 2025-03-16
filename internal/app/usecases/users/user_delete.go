package usecases

import db "firstProject/internal/infrastructure/repository"

type UserDeleteUsecase struct {
	UserRepo *db.UserRepository
}

func NewUserDeleteUsecase(userRepo *db.UserRepository) *UserDeleteUsecase {
	return &UserDeleteUsecase{UserRepo: userRepo}
}

func (u *UserDeleteUsecase) DeleteUser(id int) error {
	return u.UserRepo.DeleteUser(id)
}
