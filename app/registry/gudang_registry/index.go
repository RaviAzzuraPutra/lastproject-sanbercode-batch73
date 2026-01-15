package gudang_registry

import (
	"last-project/app/controller/gudang_controller"
	"last-project/app/repository/gudang_repository"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/gudang_service"
)

type Gudang_Module struct {
	GudangController *gudang_controller.Gudang_Controller
}

func Gudang_Registry() *Gudang_Module {

	GudangRepository := gudang_repository.NewGudangRepositoryRegistry()
	TokoRepository := toko_repository.NewTokoRepositoryResgistry()

	GudangService := gudang_service.NewGudangServiceRegistry(GudangRepository, TokoRepository)

	GudangController := gudang_controller.NewGudangControllerRegistry(GudangService)

	return &Gudang_Module{
		GudangController: GudangController,
	}

}
