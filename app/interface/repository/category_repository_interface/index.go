package category_repository_interface

import "last-project/app/models"

type Category_Repository_Interface interface {
	Create(category *models.Category) error
	GetByIdToko(IDToko string) ([]models.Category, error)
	GetByIdAndIdToko(ID string, IDToko string) (*models.Category, error)
	Update(ID string, IDToko string, category *models.Category) error
	Delete(ID string, IDToko string) error
}
