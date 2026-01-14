package user_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type User_Repository struct {
	DB *gorm.DB
}

func NewUserRepositoryRegistry() *User_Repository {
	return &User_Repository{
		DB: database.DB,
	}
}

func (repo *User_Repository) GetById(ID string) (*models.User, error) {
	var user *models.User

	errGet := repo.DB.Table("user").Where("id = ?", ID).First(&user).Error

	return user, errGet
}

func (repo *User_Repository) Update(ID string, user *models.User) error {
	errUpdate := repo.DB.Table("user").Where("id = ?", ID).Updates(user).Error

	return errUpdate
}

func (repo *User_Repository) Delete(ID string) error {

	var user *models.User

	errDelete := repo.DB.Table("user").Unscoped().Where("id = ?", ID).Delete(&user).Error

	return errDelete
}
