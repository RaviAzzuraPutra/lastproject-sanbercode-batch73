package auth_repository_interface

import "last-project/app/models"

type Auth_Repository_Interface interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
	IsEmailExist(email string) error
	IsPhoneExist(phone string) error
}
