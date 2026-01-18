package trx_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/trx_request"
)

type Trx_Service_Interface interface {
	Create(request *trx_request.Trx_Log_Request, IDGudang string, IDBarang string) (*models.Trx_Log, error)
	GetByIdBarang(IDBarang string, IDGudang string) ([]models.Trx_Log, error)
}
