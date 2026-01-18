package database

import (
	"fmt"
	"last-project/app/config/db_config"
	"last-project/app/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	var errorConnect error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().DB_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD,
		db_config.DB_Config().DB_NAME, db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SLLMODE, db_config.DB_Config().DB_TIMEZONE)

	DB, errorConnect = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errorConnect != nil {
		return fmt.Errorf("An error occurred while attempting to connect to the database. " + errorConnect.Error())
	}

	DB.Exec("CREATE TYPE trx_type AS ENUM ('in', 'out');")

	errMigrate := DB.AutoMigrate(&models.Barang{}, &models.Category{}, &models.Gudang{}, &models.Smart_Log{}, &models.Toko{}, &models.Trx_Log{}, &models.User{})

	if errMigrate != nil {
		return fmt.Errorf("Failed to Migrate Database " + errMigrate.Error())
	}

	log.Println("Successfully connected to the database ✅")

	return nil
}
