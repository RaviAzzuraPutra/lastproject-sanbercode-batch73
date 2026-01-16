package category_service

import (
	"last-project/app/helper"
	"last-project/app/interface/repository/category_repository_interface"
	"last-project/app/interface/repository/toko_repository_interface"
	"last-project/app/models"
	"last-project/app/request/category_request"
	"time"
)

type Category_Service struct {
	repository category_repository_interface.Category_Repository_Interface
	toko       toko_repository_interface.Toko_Repository_Interface
}

func NewCategoryServiceRegisry(repository category_repository_interface.Category_Repository_Interface,
	toko toko_repository_interface.Toko_Repository_Interface) *Category_Service {
	return &Category_Service{
		repository: repository,
		toko:       toko,
	}
}

func (s *Category_Service) Create(request *category_request.Category_Request, IDUser string) (*models.Category, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Description == nil || *request.Description == "" {
		return nil, helper.NewBadRequest("Description cannot be empty")
	}

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	category := &models.Category{
		Name:        request.Name,
		Description: request.Description,
		TokoID:      toko.ID,
	}

	errCreate := s.repository.Create(category)

	if errCreate != nil {
		return nil, helper.NewInternalServerError("An error occurred while adding category data. " + errCreate.Error())
	}

	return category, nil
}

func (s *Category_Service) GetByIdToko(IDUser string) ([]models.Category, error) {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	category, errGet := s.repository.GetByIdToko(*toko.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if category == nil {
		return nil, helper.NewNotFound("Category Not Found")
	}

	return category, nil

}

func (s *Category_Service) GetByIdAndIdToko(ID string, IDUser string) (*models.Category, error) {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	category, errGet := s.repository.GetByIdAndIdToko(ID, *toko.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if category == nil {
		return nil, helper.NewNotFound("Category Not Found")
	}

	return category, nil

}

func (s *Category_Service) Update(ID string, IDUser string, request *category_request.Category_Request) (*models.Category, error) {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	category, errGet := s.repository.GetByIdAndIdToko(ID, *toko.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if category == nil {
		return nil, helper.NewNotFound("Category Not Found")
	}

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Description == nil || *request.Description == "" {
		return nil, helper.NewBadRequest("Description cannot be empty")
	}

	category.Name = request.Name
	category.Description = request.Description
	category.UpdatedAt = time.Now()

	errUpdate := s.repository.Update(ID, *toko.ID, category)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while update category data. " + errUpdate.Error())
	}

	return category, nil

}

func (s *Category_Service) Delete(ID string, IDUser string) error {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	category, errGet := s.repository.GetByIdAndIdToko(ID, *toko.ID)

	if errGet != nil {
		return helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if category == nil {
		return helper.NewNotFound("Category Not Found")
	}

	errDelete := s.repository.Delete(ID, *toko.ID)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while delete category data. " + errDelete.Error())
	}

	return nil

}
