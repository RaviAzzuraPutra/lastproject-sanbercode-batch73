package config

import (
	"last-project/app/config/cloudinary_config"
	"last-project/app/config/db_config"
	"last-project/app/config/gemini_config"
	"last-project/app/config/jwt_config"
	"last-project/app/config/port_config"
)

func AppConfig() {
	port_config.Port_Config()
	gemini_config.Gemini_Config()
	db_config.DB_Config()
	cloudinary_config.Cloudinary_Config()
	jwt_config.JWT_Config()
}
