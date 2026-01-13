package auth_router

import (
	"last-project/app/controller/auth_controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(app *gin.Engine, AuthController *auth_controller.Auth_Controller) {
	auth := app.Group("/api/auth")

	auth.POST("/register", AuthController.Register)
	auth.POST("/login", AuthController.Login)
}
