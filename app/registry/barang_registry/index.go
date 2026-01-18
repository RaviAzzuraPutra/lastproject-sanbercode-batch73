package barang_registry

import (
	"last-project/app/controller/barang_controller"
	"last-project/app/repository/barang_repository"
	"last-project/app/repository/category_repository"
	"last-project/app/repository/gudang_repository"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/barang_service"
)

type Barang_Module struct {
	BarangController *barang_controller.Barang_Controller
}

func Barang_Registry() *Barang_Module {

	BarangRepository := barang_repository.NewBarangRepositoryRegistry()
	GudangRepository := gudang_repository.NewGudangRepositoryRegistry()
	TokoRepository := toko_repository.NewTokoRepositoryResgistry()
	CategoryRepository := category_repository.NewCategoryRepositoryRegistry()

	BarangService := barang_service.NewBarangServiceRegistry(BarangRepository, GudangRepository, CategoryRepository, TokoRepository)

	BarangController := barang_controller.NewBarangControllerRegistry(BarangService)

	return &Barang_Module{
		BarangController: BarangController,
	}

}
