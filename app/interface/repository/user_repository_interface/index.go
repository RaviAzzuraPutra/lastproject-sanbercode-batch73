package user_repository_interface

import "last-project/app/models"

type User_Repository_Interface interface {
	GetById(ID string) (*models.User, error)
	Update(ID string, user *models.User) error
	Delete(ID string) error
}
