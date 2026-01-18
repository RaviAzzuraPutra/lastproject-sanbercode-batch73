package smartlog_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type SmartLog_Repository struct {
	DB *gorm.DB
}

func NewSmartLogRepositoryRegistry() *SmartLog_Repository {
	return &SmartLog_Repository{
		DB: database.DB,
	}
}

func (repo *SmartLog_Repository) Create(smart *models.Smart_Log) error {

	errCreate := repo.DB.Table("smart_log").Create(smart).Error

	return errCreate

}
