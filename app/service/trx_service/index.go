package trx_service

import (
	"last-project/app/helper"
	"last-project/app/interface/repository/barang_repository_interface"
	"last-project/app/interface/repository/smartlog_repository_interface"
	"last-project/app/interface/repository/trx_repository_interface"
	"last-project/app/models"
	"last-project/app/request/trx_request"
	"time"
)

type Trx_Service struct {
	repository trx_repository_interface.Trx_Repository_Interface
	smartLog   smartlog_repository_interface.SmartLog_Repository_Interface
	barang     barang_repository_interface.Barang_Repository_Interface
}

func NewTrxServiceRegistry(repository trx_repository_interface.Trx_Repository_Interface,
	smartLog smartlog_repository_interface.SmartLog_Repository_Interface,
	barang barang_repository_interface.Barang_Repository_Interface) *Trx_Service {
	return &Trx_Service{
		repository: repository,
		smartLog:   smartLog,
		barang:     barang,
	}
}

func (s *Trx_Service) Create(request *trx_request.Trx_Log_Request, IDGudang string, IDBarang string) (*models.Trx_Log, error) {

	if *request.Qty < 0 || request.Qty == nil {
		return nil, helper.NewBadRequest("Quantity cannot under 0 or cannot be empty")
	}

	if request.Type == nil || *request.Type == "" {
		return nil, helper.NewBadRequest("Type cannot be empty")
	}

	barang, errBarang := s.barang.GetByIdGudangAndId(IDGudang, IDBarang)

	if errBarang != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Barang Data " + errBarang.Error())
	}

	currentStock := 0

	if barang.Stock != nil {
		currentStock = *barang.Stock
	}

	var finalStock int
	trxType := models.TrxLogType(*request.Type)

	if trxType == models.TrxIn {
		finalStock = currentStock + *request.Qty
	} else if trxType == models.TrxOut {
		if currentStock < *request.Qty {
			return nil, helper.NewBadRequest("Stock is not enough")
		}
		finalStock = currentStock - *request.Qty
	}

	needRestock := false
	if barang.Safety_stock != nil && finalStock <= *barang.Safety_stock {
		needRestock = true
	} else if finalStock > *barang.Safety_stock {
		needRestock = false
	}

	trx := &models.Trx_Log{
		BarangID: barang.ID,
		GudangID: &IDGudang,
		Qty:      request.Qty,
		Type:     trxType,
	}

	errCreate := s.repository.Create(trx)

	if errCreate != nil {
		return nil, helper.NewInternalServerError("An error occurred while create trx log. " + errCreate.Error())
	}

	barang.Stock = &finalStock
	barang.Need_restock = &needRestock
	barang.UpdatedAt = time.Now()

	errUpdate := s.barang.Update(*barang.GudangID, *barang.ID, barang)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while update barang. " + errUpdate.Error())
	}

	itemName := ""
	if barang.Name != nil {
		itemName = *barang.Name
	}

	latestStock := *barang.Stock

	_, aiInsight, errAi := helper.GeminiInsight(itemName, latestStock, string(trxType), *request.Qty)

	if errAi == nil {
		month := int(time.Now().Month())
		year := time.Now().Year()

		smartLog := &models.Smart_Log{
			BarangID:     barang.ID,
			GudangID:     &IDGudang,
			Period_month: &month,
			Period_year:  &year,
			AI_Insight:   &aiInsight,
		}

		errSmartLog := s.smartLog.Create(smartLog)

		if errSmartLog != nil {
			return nil, helper.NewInternalServerError("An error occurred while create smart log. " + errSmartLog.Error())
		}

	}

	return trx, nil

}

func (s *Trx_Service) GetByIdBarang(IDBarang string, IDGudang string) ([]models.Trx_Log, error) {

	barang, errBarang := s.barang.GetByIdGudangAndId(IDGudang, IDBarang)

	if errBarang != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Barang Data " + errBarang.Error())
	}

	trx, errGet := s.repository.GetByIdBarang(*barang.ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Retrieving Trx Data " + errGet.Error())
	}

	return trx, nil
}
