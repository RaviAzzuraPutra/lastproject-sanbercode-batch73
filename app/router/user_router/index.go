package user_router

import (
	"last-project/app/controller/user_controller"
	"last-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func User_Register(app *gin.Engine, UserController *user_controller.User_Controller) {
	user := app.Group("/api/profile", middleware.JWTMiddleware())

	user.GET("/", UserController.GetById)
	user.PUT("/update", UserController.Update)
	user.DELETE("/delete", UserController.Delete)
}
