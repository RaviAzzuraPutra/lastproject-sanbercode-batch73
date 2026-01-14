package toko_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Toko_Repository struct {
	DB *gorm.DB
}

func NewTokoRepositoryResgistry() *Toko_Repository {
	return &Toko_Repository{
		DB: database.DB,
	}
}

func (repo *Toko_Repository) GetByIdAndIdUser(ID string, IDUser string) (*models.Toko, error) {
	var toko *models.Toko

	errGet := repo.DB.Table("toko").Where("id = ? and user_id = ?", ID, IDUser).First(&toko).Error

	return toko, errGet
}

func (repo *Toko_Repository) Update(ID string, IDUser string, toko *models.Toko) error {
	errUpdate := repo.DB.Table("toko").Where("id = ? and user_id = ?", ID, IDUser).Updates(toko).Error

	return errUpdate
}
