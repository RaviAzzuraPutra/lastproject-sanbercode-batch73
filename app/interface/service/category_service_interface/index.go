package category_service_interface

import (
	"last-project/app/models"
	"last-project/app/request/category_request"
)

type Category_Service_Interface interface {
	Create(request *category_request.Category_Request, IDUser string) (*models.Category, error)
	GetByIdToko(IDUser string) ([]models.Category, error)
	GetByIdAndIdToko(ID string, IDUser string) (*models.Category, error)
	Update(ID string, IDUser string, request *category_request.Category_Request) (*models.Category, error)
	Delete(ID string, IDUser string) error
}
