package category_registry

import (
	"last-project/app/controller/category_controller"
	"last-project/app/repository/category_repository"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/category_service"
)

type Category_Module struct {
	CategoryController *category_controller.Category_Controller
}

func Category_Registry() *Category_Module {
	CategoryRepository := category_repository.NewCategoryRepositoryRegistry()
	TokoRepository := toko_repository.NewTokoRepositoryResgistry()

	CategoryService := category_service.NewCategoryServiceRegisry(CategoryRepository, TokoRepository)

	CategoryController := category_controller.NewCategoryControllerRegistry(CategoryService)

	return &Category_Module{
		CategoryController: CategoryController,
	}
}
