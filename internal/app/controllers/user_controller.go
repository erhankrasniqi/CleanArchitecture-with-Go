package controllers

import (
	usecases "firstProject/internal/app/usecases/users"
	domain "firstProject/internal/domain/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase       *usecases.UserUsecase
	UserByIDUsecase   *usecases.UserByIDUsecase
	UserUpdateUsecase *usecases.UserUpdateUsecase
	UserDeleteUsecase *usecases.UserDeleteUsecase
}

// NewUserController krijon një instancë të UserController
func NewUserController(userUsecase *usecases.UserUsecase, userByIDUsecase *usecases.UserByIDUsecase, userUpdateUsecase *usecases.UserUpdateUsecase, useUserDeleteUsecase *usecases.UserDeleteUsecase) *UserController {
	return &UserController{
		UserUsecase:       userUsecase,
		UserByIDUsecase:   userByIDUsecase,
		UserUpdateUsecase: userUpdateUsecase,
		UserDeleteUsecase: useUserDeleteUsecase,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User "OK"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Gabim gjatë marrjes së përdoruesve"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a single user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User "OK"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.UserByIDUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update user details
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body domain.UpdateUser true "Updated user details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser domain.UpdateUser
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// ✅ Thirr validimin
	if err := updatedUser.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userToUpdate := domain.User{
		Id:       id,
		UserName: updatedUser.UserName,
		Email:    updatedUser.Email,
	}

	err = uc.UserUpdateUsecase.UpdateUser(id, userToUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Kërko fshirjen e përdoruesit
	err = uc.UserDeleteUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// AddUser godoc
// @Summary Add a new user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.AddUser true "New user details"
// @Success 201 {object} domain.User "User added successfully"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /users [post]
func (uc *UserController) AddUser(c *gin.Context) {
	var newUser domain.AddUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	if err := newUser.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.UserUsecase.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
