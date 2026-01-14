package toko_registry

import (
	"last-project/app/controller/toko_controller"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/toko_service"
)

type Toko_Module struct {
	TokoController *toko_controller.Toko_Controller
}

func Toko_Registry() *Toko_Module {
	TokoRepository := toko_repository.NewTokoRepositoryResgistry()

	TokoService := toko_service.NewTokoServiceRegistry(TokoRepository)

	TokoController := toko_controller.NewTokoControllerRegistry(TokoService)

	return &Toko_Module{
		TokoController: TokoController,
	}
}
