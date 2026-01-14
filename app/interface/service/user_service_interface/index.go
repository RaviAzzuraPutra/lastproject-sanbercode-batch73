package user_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/user_request"
)

type User_Service_Interface interface {
	GetById(ID string) (*models.User, error)
	Update(request *user_request.User_Request, ID string) (*models.User, error)
	Delete(ID string) error
}
