package gudang_controller

import (
	"last-project/app/helper"
	"last-project/app/interface/service/gudang_service_interface"
	"last-project/app/request/gudang_request"

	"github.com/gin-gonic/gin"
)

type Gudang_Controller struct {
	service gudang_service_interface.Gudang_Service_Interface
}

func NewGudangControllerRegistry(service gudang_service_interface.Gudang_Service_Interface) *Gudang_Controller {
	return &Gudang_Controller{
		service: service,
	}
}

func (c *Gudang_Controller) Create(ctx *gin.Context) {
	request := new(gudang_request.Gudang_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
		return
	}

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userId := value.(string)

	gudang, errCreate := c.service.Create(request, userId)

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
		"Message": "Success Create Gudang",
		"Data":    gudang,
	})

}

func (c *Gudang_Controller) GetByIdToko(ctx *gin.Context) {

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userId := value.(string)

	gudang, errGet := c.service.GetByIdToko(userId)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
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
		"Message": "Success Get Gudang",
		"Data":    gudang,
	})
}

func (c *Gudang_Controller) GetByIdAndIdToko(ctx *gin.Context) {

	ID := ctx.Param("id")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userId := value.(string)

	gudang, errGet := c.service.GetByIdAndIdToko(ID, userId)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
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
		"Message": "Success Get Gudang",
		"Data":    gudang,
	})

}

func (c *Gudang_Controller) Update(ctx *gin.Context) {

	request := new(gudang_request.Gudang_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
		return
	}

	ID := ctx.Param("id")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userId := value.(string)

	gudang, errUpdate := c.service.Update(ID, userId, request)

	if errUpdate != nil {
		if appError, ok := errUpdate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errUpdate.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Update Gudang",
		"Data":    gudang,
	})
}

func (c *Gudang_Controller) Delete(ctx *gin.Context) {

	ID := ctx.Param("id")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userId := value.(string)

	errDelete := c.service.Delete(ID, userId)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Delete Gudang",
	})
}
