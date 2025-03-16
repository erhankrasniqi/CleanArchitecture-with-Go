package main

// @title My API
// @version 1.0
// @description This is a sample API documentation
// @host localhost:8080
// @BasePath /
import (
	"fmt"

	"firstProject/internal/app/controllers"
	addressUsecases "firstProject/internal/app/usecases/address"
	userUsecases "firstProject/internal/app/usecases/users"
	"firstProject/internal/infrastructure/db"
	repo "firstProject/internal/infrastructure/repository"

	"firstProject/internal/routes"

	_ "firstProject/docs" // Swagger docs import

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectToDB()
	dbInstance := db.GetDB()

	userRepository := repo.NewUserRepository(dbInstance)
	addressRepository := repo.NewAddressRepository(dbInstance)

	userUsecase := userUsecases.NewUserUsecase(userRepository)
	userByIDUsecase := userUsecases.NewUserByIDUsecase(userRepository)
	userUpdateUsecase := userUsecases.NewUserUpdateUsecase(userRepository)
	userDeleteUsecase := userUsecases.NewUserDeleteUsecase(userRepository)

	addressUsecase := addressUsecases.NewAddressUsecases(addressRepository)

	userController := controllers.NewUserController(userUsecase, userByIDUsecase, userUpdateUsecase, userDeleteUsecase)
	addressController := controllers.NewAddressController(addressUsecase)

	r := gin.Default()

	routes.SetupRoutes(r, userController, addressController)

	fmt.Println("Serveri starting with port 8080...")
	r.Run(":8080")
}
