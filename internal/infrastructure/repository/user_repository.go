package db

import (
	"firstProject/internal/domain/users"
	domain "firstProject/internal/domain/users"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(id int, updatedUser domain.User) error {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return err
	}

	updates := map[string]interface{}{}
	if updatedUser.UserName != "" {
		updates["username"] = updatedUser.UserName
	}
	if updatedUser.Email != "" {
		updates["email"] = updatedUser.Email
	}

	fmt.Printf("Updating user ID %d with values: %+v\n", id, updates)

	return r.DB.Model(&user).Updates(updates).Error
}

func (r *UserRepository) DeleteUser(id int) error {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return err
	}

	return r.DB.Delete(&user).Error
}

func (r *UserRepository) AddUser(newUser users.User) (users.User, error) {
	if err := r.DB.Create(&newUser).Error; err != nil {
		fmt.Println("Error while inserting user:", err)
		return users.User{}, err
	}
	return newUser, nil
}
