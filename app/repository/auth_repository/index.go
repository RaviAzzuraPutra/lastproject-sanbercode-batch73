package auth_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Auth_Repository struct {
	DB *gorm.DB
}

func NewAuthRepositoryRegistry() *Auth_Repository {
	return &Auth_Repository{
		DB: database.DB,
	}
}

func (repo *Auth_Repository) Register(user *models.User) error {
	errRegister := repo.DB.Table("user").Create(user).Error

	return errRegister
}

func (repo *Auth_Repository) Login(email string) (*models.User, error) {

	var user models.User

	errLogin := repo.DB.Table("user").Where("email = ?", email).First(&user).Error

	return &user, errLogin
}

func (repo *Auth_Repository) IsEmailExist(email string) error {

	var user models.User

	errFind := repo.DB.Table("user").Where("email = ?", email).First(&user).Error

	return errFind
}

func (repo *Auth_Repository) IsPhoneExist(phone string) error {

	var user models.User

	errFind := repo.DB.Table("user").Where("no_telp = ?", phone).First(&user).Error

	return errFind
}
