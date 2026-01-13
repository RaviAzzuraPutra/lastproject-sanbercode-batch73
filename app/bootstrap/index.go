package bootstrap

import (
	"last-project/app/config"
	"last-project/app/config/port_config"
	"last-project/app/database"
	"last-project/app/registry/auth_registry"
	"last-project/app/router/auth_router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("An error occurred while loading ENV!!!")
	}

	config.AppConfig()

	database.Connect()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "The application is running smoothly 👍",
		})
	})

	AuthModules := auth_registry.AuthRegistry()

	auth_router.AuthRouter(app, AuthModules.Auth_Controller)

	app.Run(port_config.PORT)
}
