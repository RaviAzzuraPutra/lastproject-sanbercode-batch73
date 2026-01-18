package trx_repository_interface

import "last-project/app/models"

type Trx_Repository_Interface interface {
	Create(trx *models.Trx_Log) error
	GetByIdBarang(IDBarang string) ([]models.Trx_Log, error)
}
