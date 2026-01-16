package category_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Category_Repository struct {
	DB *gorm.DB
}

func NewCategoryRepositoryRegistry() *Category_Repository {
	return &Category_Repository{
		DB: database.DB,
	}
}

func (repo *Category_Repository) Create(category *models.Category) error {

	errCreate := repo.DB.Table("category").Create(category).Error

	return errCreate

}

func (repo *Category_Repository) GetByIdToko(IDToko string) ([]models.Category, error) {

	var category []models.Category

	errGet := repo.DB.Table("category").Where("toko_id = ?", IDToko).Find(&category).Error

	return category, errGet

}

func (repo *Category_Repository) GetByIdAndIdToko(ID string, IDToko string) (*models.Category, error) {

	var category *models.Category

	errGet := repo.DB.Table("category").Where("id = ? AND toko_id = ?", ID, IDToko).First(&category).Error

	return category, errGet

}

func (repo *Category_Repository) Update(ID string, IDToko string, category *models.Category) error {

	errUpdate := repo.DB.Table("category").Where("id = ? AND toko_id = ?", ID, IDToko).Updates(category).Error

	return errUpdate

}

func (repo *Category_Repository) Delete(ID string, IDToko string) error {

	var category *models.Category

	errDelete := repo.DB.Table("category").Unscoped().Where("id = ? AND toko_id = ?", ID, IDToko).Delete(&category).Error

	return errDelete

}
