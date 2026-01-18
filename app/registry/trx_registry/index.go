package trx_registry

import (
	"last-project/app/controller/trx_controller"
	"last-project/app/repository/barang_repository"
	"last-project/app/repository/smartlog_repository"
	"last-project/app/repository/trx_repository"
	"last-project/app/service/trx_service"
)

type Trx_Module struct {
	TrxController *trx_controller.Trx_Controller
}

func Trx_Registry() *Trx_Module {
	TrxRepository := trx_repository.NewTrxRepositoryRegistry()
	SmartLogRepository := smartlog_repository.NewSmartLogRepositoryRegistry()
	BarangRepository := barang_repository.NewBarangRepositoryRegistry()

	TrxService := trx_service.NewTrxServiceRegistry(TrxRepository, SmartLogRepository, BarangRepository)

	TrxController := trx_controller.NewTrxControllerRegistry(TrxService)

	return &Trx_Module{
		TrxController: TrxController,
	}
}
