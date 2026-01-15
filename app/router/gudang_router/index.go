package gudang_router

import (
	"last-project/app/controller/gudang_controller"
	"last-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func GudangRouter(app *gin.Engine, GudangController *gudang_controller.Gudang_Controller) {
	gudang := app.Group("/api/gudang", middleware.JWTMiddleware())

	gudang.POST("/create", GudangController.Create)
	gudang.GET("/", GudangController.GetByIdToko)
	gudang.GET("/:id", GudangController.GetByIdAndIdToko)
	gudang.PUT("/update/:id", GudangController.Update)
	gudang.DELETE("/delete/:id", GudangController.Delete)
}
