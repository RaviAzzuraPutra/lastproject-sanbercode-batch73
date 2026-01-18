package trx_router

import (
	"last-project/app/controller/trx_controller"
	"last-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func Trx_Router(app *gin.Engine, TrxController *trx_controller.Trx_Controller) {

	trx := app.Group("/api/trx/:id_gudang/barang/:id_barang", middleware.JWTMiddleware())

	trx.POST("/create", TrxController.Create)
	trx.GET("/detail", TrxController.GetByIdAndIdBarang)

}
