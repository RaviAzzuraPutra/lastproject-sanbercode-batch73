package toko_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/toko_request"
)

type Toko_Service_Interface interface {
	GetByIdAndIdUser(IDUser string) (*models.Toko, error)
	UpdateToko(ID string, IDUser string, request *toko_request.Toko_Request) (*models.Toko, error)
}
