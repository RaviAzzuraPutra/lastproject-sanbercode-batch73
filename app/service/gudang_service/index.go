package gudang_service

import (
	"last-project/app/helper"
	"last-project/app/interface/repository/gudang_repository_interface"
	"last-project/app/interface/repository/toko_repository_interface"
	"last-project/app/models"
	"last-project/app/request/gudang_request"
	"time"
)

type Gudang_Service struct {
	repository gudang_repository_interface.Gudang_Repository_Interface
	toko       toko_repository_interface.Toko_Repository_Interface
}

func NewGudangServiceRegistry(repository gudang_repository_interface.Gudang_Repository_Interface,
	toko toko_repository_interface.Toko_Repository_Interface) *Gudang_Service {
	return &Gudang_Service{
		repository: repository,
		toko:       toko,
	}
}

func (s *Gudang_Service) Create(request *gudang_request.Gudang_Request, IDUser string) (*models.Gudang, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Address == nil || *request.Address == "" {
		return nil, helper.NewBadRequest("Address cannot be empty")
	}

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	gudang := &models.Gudang{
		TokoID:  toko.ID,
		Name:    request.Name,
		Address: request.Address,
	}

	errCreate := s.repository.Create(gudang)

	if errCreate != nil {
		return nil, helper.NewInternalServerError("An error occurred while adding gudang data. " + errCreate.Error())
	}

	return gudang, nil
}

func (s *Gudang_Service) GetByIdToko(IDUser string) ([]models.Gudang, error) {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	gudang, errGet := s.repository.GetByIdToko(*toko.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Gudang Data " + errGet.Error())
	}

	if gudang == nil {
		return nil, helper.NewNotFound("Gudang not found")
	}

	return gudang, nil

}

func (s *Gudang_Service) GetByIdAndIdToko(ID string, IDUser string) (*models.Gudang, error) {
	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	gudang, errGet := s.repository.GetByIdAndByIdToko(ID, *toko.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Gudang Data " + errGet.Error())
	}

	if gudang == nil {
		return nil, helper.NewNotFound("Gudang not found")
	}

	return gudang, nil
}

func (s *Gudang_Service) Update(ID string, IDUser string, request *gudang_request.Gudang_Request) (*models.Gudang, error) {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	gudang, errGet := s.repository.GetByIdAndByIdToko(ID, *toko.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Gudang Data " + errGet.Error())
	}

	if gudang == nil {
		return nil, helper.NewNotFound("Gudang not found")
	}

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Address == nil || *request.Address == "" {
		return nil, helper.NewBadRequest("Address cannot be empty")
	}

	gudang.Name = request.Name
	gudang.Address = request.Address
	gudang.UpdatedAt = time.Now()

	errUpdate := s.repository.Update(*gudang.ID, *toko.ID, gudang)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while updated gudang data. " + errUpdate.Error())
	}

	return gudang, nil

}

func (s *Gudang_Service) Delete(ID string, IDUser string) error {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return helper.NewInternalServerError("An Error Occurred While Retrieving Toko Data " + errToko.Error())
	}

	gudang, errGet := s.repository.GetByIdAndByIdToko(ID, *toko.ID)

	if errGet != nil {
		return helper.NewInternalServerError("An Error Occurred While Retrieving Gudang Data " + errGet.Error())
	}

	if gudang == nil {
		return helper.NewNotFound("Gudang not found")
	}

	errDelete := s.repository.Delete(*gudang.ID, *toko.ID)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while delete gudang data. " + errDelete.Error())
	}

	return nil
}
