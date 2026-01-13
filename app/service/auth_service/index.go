package auth_service

import (
	"errors"
	"last-project/app/helper"
	"last-project/app/interface/repository/auth_repository_interface"
	"last-project/app/models"
	"last-project/app/request/auth_request"

	"gorm.io/gorm"
)

type Auth_Service struct {
	repository auth_repository_interface.Auth_Repository_Interface
}

func NewAuthServiceRegistry(repository auth_repository_interface.Auth_Repository_Interface) *Auth_Service {
	return &Auth_Service{
		repository: repository,
	}
}

func (s *Auth_Service) Register(request *auth_request.Register_Request) (*models.User, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name Cannot Be Empty!")
	}

	if request.Email == nil || *request.Email == "" {
		return nil, helper.NewBadRequest("Email Cannot Be Empty!")
	}

	if request.No_Telp == nil || *request.No_Telp == "" {
		return nil, helper.NewBadRequest("Phone Number Cannot Be Empty!")
	}

	if request.TokoName == nil || *request.TokoName == "" {
		return nil, helper.NewBadRequest("Toko Name Cannot Be Empty!")
	}

	if request.Password == nil || *request.Password == "" {
		return nil, helper.NewBadRequest("Password Cannot Be Empty!")
	}

	HashPassword, errHash := helper.HashPassword(*request.Password)

	if errHash != nil {
		return nil, helper.NewInternalServerError("An error occurred while hashing the password. " + errHash.Error())
	}

	errExistingEmail := s.repository.IsEmailExist(*request.Email)

	if errExistingEmail != nil && !errors.Is(errExistingEmail, gorm.ErrRecordNotFound) {
		return nil, helper.NewInternalServerError("Database Error While Checking Email " + errExistingEmail.Error())
	}

	if errExistingEmail == nil {
		return nil, helper.NewBadRequest("Email Already Registered No Duplicates Allowed")
	}

	errExistingPhone := s.repository.IsPhoneExist(*request.No_Telp)

	if errExistingPhone != nil && !errors.Is(errExistingPhone, gorm.ErrRecordNotFound) {
		return nil, helper.NewInternalServerError("Database Error While Checking Phone " + errExistingPhone.Error())
	}

	if errExistingPhone == nil {
		return nil, helper.NewBadRequest("Phone Number Already Registered No Duplicates Allowed")
	}

	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: &HashPassword,
		No_Telp:  request.No_Telp,
		Toko: &models.Toko{
			Name: request.TokoName,
		},
	}

	errRegister := s.repository.Register(user)

	if errRegister != nil {
		return nil, helper.NewInternalServerError("An error occurred during registration. " + errRegister.Error())
	}

	return user, nil
}

func (s *Auth_Service) Login(request *auth_request.Login_Request) (string, error) {

	if request.Email == nil || *request.Email == "" {
		return "", helper.NewBadRequest("Email Cannot Be Empty!")
	}

	if request.Password == nil || *request.Password == "" {
		return "", helper.NewBadRequest("Password Cannot Be Empty!")
	}

	login, errLogin := s.repository.Login(*request.Email)

	if errLogin != nil {
		return "", helper.NewInternalServerError("An error occurred while attempting to log in. " + errLogin.Error())
	}

	errCheckPassword := helper.CheckPassword(*login.Password, *request.Password)

	if errCheckPassword != nil {
		return "", helper.NewInternalServerError("Password Does Not Match " + errCheckPassword.Error())
	}

	JWT, errJWT := helper.GenerateJWT(*login.ID)

	if errJWT != nil {
		return "", helper.NewInternalServerError("An error occurred while generating the token. " + errJWT.Error())
	}

	return JWT, nil
}
