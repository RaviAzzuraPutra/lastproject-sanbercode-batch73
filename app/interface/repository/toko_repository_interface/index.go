package toko_repository_interface

import "last-project/app/models"

type Toko_Repository_Interface interface {
	GetByIdAndIdUser(ID string, IDUser string) (*models.Toko, error)
	Update(ID string, IDUser string, toko *models.Toko) error
}
