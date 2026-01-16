package category_router

import (
	"last-project/app/controller/category_controller"
	"last-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(app *gin.Engine, CategoryController *category_controller.Category_Controller) {
	category := app.Group("/api/category", middleware.JWTMiddleware())

	category.POST("/create", CategoryController.Create)
	category.GET("/", CategoryController.GetByIdToko)
	category.GET("/:id", CategoryController.GetByIdAndIdToko)
	category.PUT("/update/:id", CategoryController.Update)
	category.DELETE("/delete/:id", CategoryController.Delete)
}
