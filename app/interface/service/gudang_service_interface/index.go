package gudang_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/gudang_request"
)

type Gudang_Service_Interface interface {
	Create(request *gudang_request.Gudang_Request, IDUser string) (*models.Gudang, error)
	GetByIdToko(IDUser string) ([]models.Gudang, error)
	GetByIdAndIdToko(ID string, IDUser string) (*models.Gudang, error)
	Update(ID string, IDUser string, request *gudang_request.Gudang_Request) (*models.Gudang, error)
	Delete(ID string, IDUser string) error
}
