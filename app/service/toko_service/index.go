package toko_service

import (
	"last-project/app/helper"
	"last-project/app/interface/repository/toko_repository_interface"
	"last-project/app/models"
	"last-project/app/request/toko_request"
	"time"
)

type Toko_Service struct {
	repository toko_repository_interface.Toko_Repository_Interface
}

func NewTokoServiceRegistry(repository toko_repository_interface.Toko_Repository_Interface) *Toko_Service {
	return &Toko_Service{
		repository: repository,
	}
}

func (s *Toko_Service) GetByIdUser(IDUser string) (*models.Toko, error) {
	toko, errGet := s.repository.GetByIdUser(IDUser)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errGet.Error())
	}

	if toko == nil {
		return nil, helper.NewNotFound("Toko Not Found")
	}

	return toko, errGet
}

func (s *Toko_Service) UpdateToko(ID string, IDUser string, request *toko_request.Toko_Request) (*models.Toko, error) {

	toko, errGet := s.repository.GetByIdUser(IDUser)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errGet.Error())
	}

	if toko == nil {
		return nil, helper.NewNotFound("Toko Not Found")
	}

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Address == nil || *request.Address == "" {
		return nil, helper.NewBadRequest("Address cannot be empty")
	}

	toko.Name = request.Name
	toko.Address = request.Address
	toko.UpdatedAt = time.Now()

	errUpdate := s.repository.Update(ID, IDUser, toko)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while updating toko data. " + errUpdate.Error())
	}

	return toko, nil
}
