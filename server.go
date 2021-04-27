package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lotusntp/golang-workshop-RESTful/config"
	"github.com/lotusntp/golang-workshop-RESTful/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
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
