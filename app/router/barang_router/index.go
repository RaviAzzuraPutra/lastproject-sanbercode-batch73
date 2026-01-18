package barang_router

import (
	"last-project/app/controller/barang_controller"
	"last-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func Barang_Router(app *gin.Engine, BarangController *barang_controller.Barang_Controller) {

	barang := app.Group("/api/barang/:id_gudang", middleware.JWTMiddleware())

	barang.POST("/create", BarangController.Create)
	barang.GET("/", BarangController.GetByIdGudang)
	barang.GET("/category/:id_category", BarangController.GetByIdGudangAndIdCategory)
	barang.GET("/detail/:id", BarangController.GetByIdGudangAndId)
	barang.PUT("/update/:id", BarangController.Update)
	barang.DELETE("/delete/:id", BarangController.Delete)

}
