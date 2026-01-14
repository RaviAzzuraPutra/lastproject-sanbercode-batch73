package toko_controller

import (
	"last-project/app/helper"
	"last-project/app/interface/service/toko_service_interface"
	"last-project/app/request/toko_request"

	"github.com/gin-gonic/gin"
)

type Toko_Controller struct {
	service toko_service_interface.Toko_Service_Interface
}

func NewTokoControllerRegistry(service toko_service_interface.Toko_Service_Interface) *Toko_Controller {
	return &Toko_Controller{
		service: service,
	}
}

func (c *Toko_Controller) Get(ctx *gin.Context) {
	ID := ctx.Param("id")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	toko, errGet := c.service.GetByIdAndIdUser(ID, userID)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "An Error During get Toko Data!",
			"Error":   errGet.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Data",
		"Data":    toko,
	})
}

func (c *Toko_Controller) Update(ctx *gin.Context) {
	request := new(toko_request.Toko_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	ID := ctx.Param("id")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	toko, errUpdate := c.service.UpdateToko(ID, userID, request)

	if errUpdate != nil {
		if appError, ok := errUpdate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "An Error During Update Toko Data!",
			"Error":   errUpdate.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Update Toko",
		"Data":    toko,
	})
}
