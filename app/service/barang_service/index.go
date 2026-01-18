package barang_service

import (
	"fmt"
	"last-project/app/helper"
	"last-project/app/interface/repository/barang_repository_interface"
	"last-project/app/interface/repository/category_repository_interface"
	"last-project/app/interface/repository/gudang_repository_interface"
	"last-project/app/interface/repository/toko_repository_interface"
	"last-project/app/models"
	"last-project/app/request/barang_request"
	"time"
)

type Barang_Service struct {
	repository barang_repository_interface.Barang_Repository_Interface
	gudang     gudang_repository_interface.Gudang_Repository_Interface
	category   category_repository_interface.Category_Repository_Interface
	toko       toko_repository_interface.Toko_Repository_Interface
}

func NewBarangServiceRegistry(repository barang_repository_interface.Barang_Repository_Interface,
	gudang gudang_repository_interface.Gudang_Repository_Interface,
	category category_repository_interface.Category_Repository_Interface,
	toko toko_repository_interface.Toko_Repository_Interface) *Barang_Service {
	return &Barang_Service{
		repository: repository,
		gudang:     gudang,
		category:   category,
		toko:       toko,
	}
}

func (s *Barang_Service) Create(request *barang_request.Barang_Request, IDGudang string, IDUser string) (*models.Barang, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Sku == nil || *request.Sku == "" {
		return nil, helper.NewBadRequest("SKU cannot be empty")
	}

	if request.Image_url == nil || *request.Image_url == "" {
		return nil, helper.NewBadRequest("Image cannot be empty")
	}

	if *request.Stock < 0 || request.Sku == nil {
		return nil, helper.NewBadRequest("Stock cannot under 0 or cannot be empty")
	}

	if request.Safety_stock == nil || *request.Safety_stock < 0 {
		return nil, helper.NewBadRequest("Safety Stock cannot under 0 or cannot be empty")
	}

	if request.Lead_time_days == nil || *request.Lead_time_days < 0 {
		return nil, helper.NewBadRequest("Lead time days cannot under 0 or cannot be empty")
	}

	if request.Category_id == nil || *request.Category_id == "" {
		return nil, helper.NewBadRequest("Category cannot be empty")
	}

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errToko.Error())
	}

	category, errCategory := s.category.GetByIdAndIdToko(*request.Category_id, *toko.ID)

	if errCategory != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errCategory.Error())
	}

	folderName := fmt.Sprintf("toko_%s/gudang_%s", *toko.ID, IDGudang)

	uniqueFileName := fmt.Sprintf("%s-%d", *request.Name, time.Now().UnixNano())

	cloudinaryUpload, errUpload := helper.UploadFotoToCloudinary(*request.Image_url, folderName, uniqueFileName)

	if errUpload != nil {
		return nil, helper.NewInternalServerError("An error occurred while uploading the image. " + errUpload.Error())
	}

	needRestock := *request.Stock <= *request.Safety_stock

	barang := &models.Barang{
		Name:           request.Name,
		Sku:            request.Sku,
		Image_url:      &cloudinaryUpload,
		Stock:          request.Stock,
		Safety_stock:   request.Safety_stock,
		Need_restock:   &needRestock,
		Lead_time_days: request.Lead_time_days,
		CategoryID:     category.ID,
		GudangID:       &IDGudang,
	}

	errCreate := s.repository.Create(barang)

	if errCreate != nil {
		return nil, helper.NewInternalServerError("An error occurred while create barang. " + errCreate.Error())
	}

	return barang, nil

}

func (s *Barang_Service) GetByIdGudang(IDGudang string) ([]models.Barang, error) {

	barang, errGet := s.repository.GetByIdGudang(IDGudang)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if barang == nil {
		return nil, helper.NewNotFound("Barang Not Found")
	}

	return barang, nil

}

func (s *Barang_Service) GetByIdGudangAndIdCategory(GudangID string, CategoryID string, IDUser string) ([]models.Barang, error) {

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errToko.Error())
	}

	category, errCategory := s.category.GetByIdAndIdToko(CategoryID, *toko.ID)

	if errCategory != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errCategory.Error())
	}

	barang, errGet := s.repository.GetByIdGudangAndIdCategory(GudangID, *category.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if barang == nil {
		return nil, helper.NewNotFound("Barang Not Found")
	}

	return barang, nil
}

func (s *Barang_Service) GetByIdGudangAndId(GudangID string, ID string) (*models.Barang, error) {

	barang, errGet := s.repository.GetByIdGudangAndId(GudangID, ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if barang == nil {
		return nil, helper.NewNotFound("Barang Not Found")
	}

	return barang, nil

}

func (s *Barang_Service) Update(request *barang_request.Barang_Request, GudangID string, IDUser string, ID string) (*models.Barang, error) {
	barang, errGet := s.repository.GetByIdGudangAndId(GudangID, ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if barang == nil {
		return nil, helper.NewNotFound("Barang Not Found")
	}

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name cannot be empty")
	}

	if request.Sku == nil || *request.Sku == "" {
		return nil, helper.NewBadRequest("SKU cannot be empty")
	}

	if request.Sku == nil || *request.Stock < 0 {
		return nil, helper.NewBadRequest("Stock cannot under 0 or cannot be empty")
	}

	if request.Safety_stock == nil || *request.Safety_stock < 0 {
		return nil, helper.NewBadRequest("Safety Stock cannot under 0 or cannot be empty")
	}

	if *request.Lead_time_days < 0 || request.Lead_time_days == nil {
		return nil, helper.NewBadRequest("Lead time days cannot under 0 or cannot be empty")
	}

	if request.Category_id == nil || *request.Category_id == "" {
		return nil, helper.NewBadRequest("Category cannot be empty")
	}

	toko, errToko := s.toko.GetByIdUser(IDUser)

	if errToko != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Store Data " + errToko.Error())
	}

	category, errCategory := s.category.GetByIdAndIdToko(*request.Category_id, *toko.ID)

	if errCategory != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errCategory.Error())
	}

	finalImageUrl := *barang.Image_url

	if request.Image_url != nil && *request.Image_url != "" {
		folderName := fmt.Sprintf("toko_%s/gudang_%s", *toko.ID, GudangID)
		uniqueFileName := fmt.Sprintf("%s-%d", *request.Name, time.Now().UnixNano())

		cloudinaryUpload, errUpload := helper.UploadFotoToCloudinary(*request.Image_url, folderName, uniqueFileName)
		if errUpload != nil {
			return nil, helper.NewInternalServerError("Upload error: " + errUpload.Error())
		}
		finalImageUrl = cloudinaryUpload
	}

	needRestock := *request.Stock <= *request.Safety_stock
	categoryID := category.ID

	barang.Name = request.Name
	barang.Sku = request.Sku
	barang.Image_url = &finalImageUrl
	barang.Stock = request.Stock
	barang.Need_restock = &needRestock
	barang.Lead_time_days = request.Lead_time_days
	barang.UpdatedAt = time.Now()
	barang.CategoryID = categoryID

	errUpdate := s.repository.Update(GudangID, ID, barang)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while update barang. " + errUpdate.Error())
	}

	return barang, nil
}

func (s *Barang_Service) Delete(GudangID string, ID string) error {

	barang, errGet := s.repository.GetByIdGudangAndId(GudangID, ID)

	if errGet != nil {
		return helper.NewInternalServerError("An Error Occurred While Retrieving Category Data " + errGet.Error())
	}

	if barang == nil {
		return helper.NewNotFound("Barang Not Found")
	}

	errDelete := s.repository.Delete(GudangID, ID)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while delete barang. " + errDelete.Error())
	}

	return nil

}
