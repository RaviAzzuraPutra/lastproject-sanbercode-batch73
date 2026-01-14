package user_registry

import (
	"last-project/app/controller/user_controller"
	"last-project/app/repository/user_repository"
	"last-project/app/service/user_service"
)

type User_Module struct {
	UserController *user_controller.User_Controller
}

func User_Registry() *User_Module {
	UserRepository := user_repository.NewUserRepositoryRegistry()

	UserService := user_service.NewUserServiceRegistry(UserRepository)

	UserController := user_controller.NewUserControllerRegistry(UserService)

	return &User_Module{
		UserController: UserController,
	}
}
