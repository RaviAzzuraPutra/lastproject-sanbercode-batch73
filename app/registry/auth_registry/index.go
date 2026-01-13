package auth_registry

import (
	"last-project/app/controller/auth_controller"
	"last-project/app/repository/auth_repository"
	"last-project/app/service/auth_service"
)

type Auth_Module struct {
	Auth_Controller *auth_controller.Auth_Controller
}

func AuthRegistry() *Auth_Module {
	AuthRepository := auth_repository.NewAuthRepositoryRegistry()

	AuthService := auth_service.NewAuthServiceRegistry(AuthRepository)

	AuthController := auth_controller.NewAuthControllerRegistry(AuthService)

	return &Auth_Module{
		Auth_Controller: AuthController,
	}
}
