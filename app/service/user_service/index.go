package user_service

import (
	"last-project/app/helper"
	"last-project/app/interface/repository/user_repository_interface"
	"last-project/app/models"
	"last-project/app/request/user_request"
)

type User_Service struct {
	repository user_repository_interface.User_Repository_Interface
}

func NewUserServiceRegistry(repository user_repository_interface.User_Repository_Interface) *User_Service {
	return &User_Service{
		repository: repository,
	}
}

func (s *User_Service) GetById(ID string) (*models.User, error) {
	user, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errGet.Error())
	}

	if user == nil {
		return nil, helper.NewNotFound("User Not Found")
	}

	return user, nil
}

func (s *User_Service) Update(request *user_request.User_Request, ID string) (*models.User, error) {
	user, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errGet.Error())
	}

	if user == nil {
		return nil, helper.NewNotFound("User Not Found")
	}

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name Cannot Be Empty!")
	}

	if request.Email == nil || *request.Email == "" {
		return nil, helper.NewBadRequest("Email Cannot Be Empty!")
	}

	if request.No_Telp == nil || *request.No_Telp == "" {
		return nil, helper.NewBadRequest("Phone Number Cannot Be Empty!")
	}

	if request.Password == nil || *request.Password == "" {
		return nil, helper.NewBadRequest("Password Cannot Be Empty!")
	}

	HashPassword, errHash := helper.HashPassword(*request.Password)

	if errHash != nil {
		return nil, helper.NewInternalServerError("An error occurred while hashing the password. " + errHash.Error())
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = &HashPassword
	user.No_Telp = request.No_Telp

	errUpdate := s.repository.Update(ID, user)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while update data. " + errUpdate.Error())
	}

	return user, nil
}

func (s *User_Service) Delete(ID string) error {
	user, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errGet.Error())
	}

	if user == nil {
		return helper.NewNotFound("User Not Found")
	}

	errDelete := s.repository.Delete(ID)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while update data. " + errDelete.Error())
	}

	return nil
}
