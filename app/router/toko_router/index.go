package toko_router

import (
	"last-project/app/controller/toko_controller"
	"last-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func TokoRouter(app *gin.Engine, TokoController *toko_controller.Toko_Controller) {
	toko := app.Group("/api/toko", middleware.JWTMiddleware())

	toko.GET("/", TokoController.Get)
	toko.PUT("/update/:id", TokoController.Update)
}
