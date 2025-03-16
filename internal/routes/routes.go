package routes

import (
	_ "firstProject/docs" // Swagger docs import
	"firstProject/internal/app/controllers"

	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, addressController *controllers.AddressController) {
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.POST("/users", userController.AddUser)

	r.POST("/addresses", addressController.AddAddress)
}
