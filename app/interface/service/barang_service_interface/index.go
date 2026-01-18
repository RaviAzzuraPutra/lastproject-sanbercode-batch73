package barang_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/barang_request"
)

type Barang_Service_Interface interface {
	Create(request *barang_request.Barang_Request, IDGudang string, IDUser string) (*models.Barang, error)
	GetByIdGudang(IDGudang string) ([]models.Barang, error)
	GetByIdGudangAndIdCategory(GudangID string, CategoryID string, IDUser string) ([]models.Barang, error)
	GetByIdGudangAndId(GudangID string, ID string) (*models.Barang, error)
	Update(request *barang_request.Barang_Request, GudangID string, IDUser string, ID string) (*models.Barang, error)
	Delete(GudangID string, ID string) error
}
