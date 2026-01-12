package bootstrap

import (
	"last-project/app/config"
	"last-project/app/config/port_config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("An error occurred while loading ENV!!!")
	}

	config.AppConfig()

	app := gin.Default()

	app.Run(":" + port_config.PORT)
}
