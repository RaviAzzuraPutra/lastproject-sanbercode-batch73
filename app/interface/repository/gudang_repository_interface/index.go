package gudang_repository_interface

import "last-project/app/models"

type Gudang_Repository_Interface interface {
	Create(gudang *models.Gudang) error
	GetByIdToko(IDToko string) ([]models.Gudang, error)
	GetByIdAndByIdToko(ID string, IDToko string) (*models.Gudang, error)
	Update(ID string, IDToko string, gudang *models.Gudang) error
	Delete(ID string, IDToko string) error
}
