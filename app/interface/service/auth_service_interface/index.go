package auth_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/auth_request"
)

type Auth_Service_Interface interface {
	Register(request *auth_request.Register_Request) (*models.User, error)
	Login(request *auth_request.Login_Request) (string, error)
}
