package auth_controller

import (
	"last-project/app/helper"
	"last-project/app/interface/service/auth_service_interface"
	"last-project/app/request/auth_request"

	"github.com/gin-gonic/gin"
)

type Auth_Controller struct {
	service auth_service_interface.Auth_Service_Interface
}

func NewAuthControllerRegistry(service auth_service_interface.Auth_Service_Interface) *Auth_Controller {
	return &Auth_Controller{
		service: service,
	}
}

func (c *Auth_Controller) Register(ctx *gin.Context) {
	request := new(auth_request.Register_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	register, errRegister := c.service.Register(request)

	if errRegister != nil {
		if appErr, ok := errRegister.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Message": appErr.Message,
			})
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errRegister.Error(),
		})

		return
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Register",
		"Data":    register,
	})
}

func (c *Auth_Controller) Login(ctx *gin.Context) {
	request := new(auth_request.Login_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	login, errLogin := c.service.Login(request)

	if errLogin != nil {
		if appErr, ok := errLogin.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Message": appErr.Message,
			})
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errLogin.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Login",
		"Token":   login,
	})
}
