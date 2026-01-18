package barang_controller

import (
	"fmt"
	"last-project/app/helper"
	"last-project/app/interface/service/barang_service_interface"
	"last-project/app/request/barang_request"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Barang_Controller struct {
	service barang_service_interface.Barang_Service_Interface
}

func NewBarangControllerRegistry(service barang_service_interface.Barang_Service_Interface) *Barang_Controller {
	return &Barang_Controller{
		service: service,
	}
}

func (c *Barang_Controller) Create(ctx *gin.Context) {

	request := new(barang_request.Barang_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
		return
	}

	file, errFoto := ctx.FormFile("image_url")
	if errFoto != nil {
		ctx.JSON(400, gin.H{
			"Message": "Image file is required",
			"Error":   errFoto.Error(),
		})
		return
	}

	tempPath := fmt.Sprintf("temp/%s", file.Filename)
	if err := ctx.SaveUploadedFile(file, tempPath); err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Failed to save uploaded file",
			"Error":   err.Error(),
		})
		return
	}

	defer os.Remove(tempPath)

	request.Image_url = &tempPath

	ID_Gudang := ctx.Param("id_gudang")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	barang, errCreate := c.service.Create(request, ID_Gudang, userID)

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
		"Message": "Success Input Barang",
		"Data":    barang,
	})

}

func (c *Barang_Controller) GetByIdGudang(ctx *gin.Context) {

	ID_Gudang := ctx.Param("id_gudang")

	barang, errGet := c.service.GetByIdGudang(ID_Gudang)

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
		"Message": "Success Get Barang",
		"Data":    barang,
	})

}

func (c *Barang_Controller) GetByIdGudangAndIdCategory(ctx *gin.Context) {

	ID_Gudang := ctx.Param("id_gudang")

	ID_Category := ctx.Param("id_category")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	barang, errGet := c.service.GetByIdGudangAndIdCategory(ID_Gudang, ID_Category, userID)

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
		"Message": "Success Get Barang",
		"Data":    barang,
	})

}

func (c *Barang_Controller) GetByIdGudangAndId(ctx *gin.Context) {

	ID_Gudang := ctx.Param("id_gudang")

	ID := ctx.Param("id")

	barang, errGet := c.service.GetByIdGudangAndId(ID_Gudang, ID)

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
		"Message": "Success Get Barang",
		"Data":    barang,
	})

}

func (c *Barang_Controller) Update(ctx *gin.Context) {

	request := new(barang_request.Barang_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
		return
	}

	file, errFoto := ctx.FormFile("image_url")

	if errFoto == nil {
		tempPath := fmt.Sprintf("temp/%d-%s", time.Now().UnixNano(), file.Filename)
		if err := ctx.SaveUploadedFile(file, tempPath); err == nil {
			request.Image_url = &tempPath
			defer os.Remove(tempPath)
		}
	}

	tempPath := fmt.Sprintf("temp/%d-%s", time.Now().UnixNano(), file.Filename)
	if err := ctx.SaveUploadedFile(file, tempPath); err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Failed to save uploaded file",
			"Error":   err.Error(),
		})
		return
	}

	defer os.Remove(tempPath)

	request.Image_url = &tempPath

	ID_Gudang := ctx.Param("id_gudang")

	ID := ctx.Param("id")

	value, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(401, gin.H{
			"Message": "Your session has expired or is invalid. Please log in again.",
		})
		return
	}

	userID := value.(string)

	barang, errUpdate := c.service.Update(request, ID_Gudang, userID, ID)

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

	ctx.JSON(200, gin.H{
		"Message": "Success Update Barang",
		"Data":    barang,
	})

}

func (c *Barang_Controller) Delete(ctx *gin.Context) {

	ID_Gudang := ctx.Param("id_gudang")

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID_Gudang, ID)

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
		"Message": "Success Delete Barang",
	})

}
