package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lotusntp/golang-workshop-RESTful/config"
	"github.com/lotusntp/golang-workshop-RESTful/controller"
	"github.com/lotusntp/golang-workshop-RESTful/repository"
	"github.com/lotusntp/golang-workshop-RESTful/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRouter := r.Group("api/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
	}
	r.Run()
}
