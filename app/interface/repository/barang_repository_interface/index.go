package barang_repository_interface

import "last-project/app/models"

type Barang_Repository_Interface interface {
	Create(barang *models.Barang) error
	GetByIdGudang(GudangID string) ([]models.Barang, error)
	GetByIdGudangAndIdCategory(GudangID string, CategoryID string) ([]models.Barang, error)
	GetByIdGudangAndId(GudangID string, ID string) (*models.Barang, error)
	Update(GudangID string, ID string, barang *models.Barang) error
	Delete(GudangID string, ID string) error
}
