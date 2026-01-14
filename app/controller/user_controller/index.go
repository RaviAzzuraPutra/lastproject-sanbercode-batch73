package user_controller

import (
	"last-project/app/helper"
	"last-project/app/interface/service/user_service_interface"
	"last-project/app/request/user_request"

	"github.com/gin-gonic/gin"
)

type User_Controller struct {
	service user_service_interface.User_Service_Interface
}

func NewUserControllerRegistry(service user_service_interface.User_Service_Interface) *User_Controller {
	return &User_Controller{
		service: service,
	}
}

func (c *User_Controller) GetById(ctx *gin.Context) {
	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	user, errGet := c.service.GetById(userID)

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
		"Message": "Success Get User",
		"Data":    user,
	})
}

func (c *User_Controller) Update(ctx *gin.Context) {

	request := new(user_request.User_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	user, errUpdate := c.service.Update(request, userID)

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
		"Message": "Success Update Data",
		"Data":    user,
	})
}

func (c *User_Controller) Delete(ctx *gin.Context) {
	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	errDelete := c.service.Delete(userID)

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
		"Message": "Success Delete User",
	})
}
