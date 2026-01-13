package bootstrap

import (
	"last-project/app/config"
	"last-project/app/config/port_config"
	"last-project/app/database"

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

	app.Run(port_config.PORT)
}
