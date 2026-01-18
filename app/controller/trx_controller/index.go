package trx_controller

import (
	"last-project/app/helper"
	"last-project/app/interface/service/trx_service_interface"
	"last-project/app/request/trx_request"

	"github.com/gin-gonic/gin"
)

type Trx_Controller struct {
	service trx_service_interface.Trx_Service_Interface
}

func NewTrxControllerRegistry(service trx_service_interface.Trx_Service_Interface) *Trx_Controller {
	return &Trx_Controller{
		service: service,
	}
}

func (c *Trx_Controller) Create(ctx *gin.Context) {

	request := new(trx_request.Trx_Log_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
		return
	}

	ID_Gudang := ctx.Param("id_gudang")
	ID_Barang := ctx.Param("id_barang")

	trx, errCreate := c.service.Create(request, ID_Gudang, ID_Barang)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errCreate.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Create Trx Log",
		"Data":    trx,
	})

}

func (c *Trx_Controller) GetByIdAndIdBarang(ctx *gin.Context) {

	ID_Gudang := ctx.Param("id_gudang")
	ID_Barang := ctx.Param("id_barang")

	trx, errGet := c.service.GetByIdBarang(ID_Barang, ID_Gudang)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Trx Log",
		"Data":    trx,
	})

}
