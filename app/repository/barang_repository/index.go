package barang_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Barang_Repository struct {
	DB *gorm.DB
}

func NewBarangRepositoryRegistry() *Barang_Repository {
	return &Barang_Repository{
		DB: database.DB,
	}
}

func (repo *Barang_Repository) Create(barang *models.Barang) error {

	errCreate := repo.DB.Table("barang").Create(barang).Error

	return errCreate

}

func (repo *Barang_Repository) GetByIdGudang(GudangID string) ([]models.Barang, error) {

	var barang []models.Barang

	errGet := repo.DB.Table("barang").Preload("Category").Preload("Gudang").Preload("Trx_Log", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(5)
	}).
		Preload("Smart_Log", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1)
		}).Where("gudang_id = ?", GudangID).Find(&barang).Error

	return barang, errGet

}

func (repo *Barang_Repository) GetByIdGudangAndIdCategory(GudangID string, CategoryID string) ([]models.Barang, error) {

	var barang []models.Barang

	errGet := repo.DB.Table("barang").Preload("Category").Preload("Gudang").Preload("Trx_Log", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(5)
	}).
		Preload("Smart_Log", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1)
		}).Where("gudang_id = ? AND category_id = ?", GudangID, CategoryID).Find(&barang).Error

	return barang, errGet

}

func (repo *Barang_Repository) GetByIdGudangAndId(GudangID string, ID string) (*models.Barang, error) {

	var barang *models.Barang

	errGet := repo.DB.Table("barang").Preload("Category").Preload("Gudang").Preload("Trx_Log", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(5)
	}).
		Preload("Smart_Log", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1)
		}).Where("id = ? AND gudang_id = ?", ID, GudangID).First(&barang).Error

	return barang, errGet

}

func (repo *Barang_Repository) Update(GudangID string, ID string, barang *models.Barang) error {

	errUpdate := repo.DB.Table("barang").Where("id = ? AND gudang_id = ?", ID, GudangID).Updates(barang).Error

	return errUpdate

}

func (repo *Barang_Repository) Delete(GudangID string, ID string) error {

	var barang *models.Barang

	errDelete := repo.DB.Table("barang").Unscoped().Where("id = ? AND gudang_id = ?", ID, GudangID).Delete(&barang).Error

	return errDelete

}
