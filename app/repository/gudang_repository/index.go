package gudang_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Gudang_Repository struct {
	DB *gorm.DB
}

func NewGudangRepositoryRegistry() *Gudang_Repository {
	return &Gudang_Repository{
		DB: database.DB,
	}
}

func (repo *Gudang_Repository) Create(gudang *models.Gudang) error {
	errCreate := repo.DB.Table("gudang").Create(gudang).Error

	return errCreate
}

func (repo *Gudang_Repository) GetByIdToko(IDToko string) ([]models.Gudang, error) {

	var gudang []models.Gudang

	errGet := repo.DB.Table("gudang").Where("toko_id = ?", IDToko).Find(&gudang).Error

	return gudang, errGet
}

func (repo *Gudang_Repository) GetByIdAndByIdToko(ID string, IDToko string) (*models.Gudang, error) {

	var gudang *models.Gudang

	errGet := repo.DB.Table("gudang").Where("id = ? AND toko_id = ?", ID, IDToko).First(&gudang).Error

	return gudang, errGet

}

func (repo *Gudang_Repository) Update(ID string, IDToko string, gudang *models.Gudang) error {

	errUpdate := repo.DB.Table("gudang").Where("id = ? AND toko_id = ?", ID, IDToko).Updates(gudang).Error

	return errUpdate

}

func (repo *Gudang_Repository) Delete(ID string, IDToko string) error {

	var gudang *models.Gudang

	errDelete := repo.DB.Table("gudang").Unscoped().Where("id = ? AND toko_id = ?", ID, IDToko).Delete(&gudang).Error

	return errDelete
}
