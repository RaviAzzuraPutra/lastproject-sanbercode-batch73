package trx_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Trx_Repository struct {
	DB *gorm.DB
}

func NewTrxRepositoryRegistry() *Trx_Repository {
	return &Trx_Repository{
		DB: database.DB,
	}
}

func (repo *Trx_Repository) Create(trx *models.Trx_Log) error {

	errCreate := repo.DB.Table("trx_log").Create(trx).Error

	return errCreate

}

func (repo *Trx_Repository) GetByIdBarang(IDBarang string) ([]models.Trx_Log, error) {

	var trx []models.Trx_Log

	errGet := repo.DB.Table("trx_log").Where("barang_id = ?", IDBarang).Find(&trx).Error

	return trx, errGet

}
