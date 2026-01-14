package bootstrap

import (
	"last-project/app/config"
	"last-project/app/config/port_config"
	"last-project/app/database"
	"last-project/app/registry/auth_registry"
	"last-project/app/registry/toko_registry"
	"last-project/app/router/auth_router"
	"last-project/app/router/toko_router"

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
	TokoModules := toko_registry.Toko_Registry()

	auth_router.AuthRouter(app, AuthModules.Auth_Controller)
	toko_router.TokoRouter(app, TokoModules.TokoController)

	app.Run(port_config.PORT)
}
