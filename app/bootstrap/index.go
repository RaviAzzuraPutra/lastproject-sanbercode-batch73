package bootstrap

import (
	"last-project/app/config"
	"last-project/app/config/port_config"
	"last-project/app/database"
	"last-project/app/registry/auth_registry"
	"last-project/app/registry/barang_registry"
	"last-project/app/registry/category_registry"
	"last-project/app/registry/gudang_registry"
	"last-project/app/registry/toko_registry"
	"last-project/app/registry/trx_registry"
	"last-project/app/registry/user_registry"
	"last-project/app/router/auth_router"
	"last-project/app/router/barang_router"
	"last-project/app/router/category_router"
	"last-project/app/router/gudang_router"
	"last-project/app/router/toko_router"
	"last-project/app/router/trx_router"
	"last-project/app/router/user_router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	_ = godotenv.Load()

	config.AppConfig()

	database.Connect()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "The application is running smoothly 👍",
		})
	})

	AuthModules := auth_registry.AuthRegistry()
	TokoModules := toko_registry.Toko_Registry()
	UserModules := user_registry.User_Registry()
	GudangModules := gudang_registry.Gudang_Registry()
	CategoryModules := category_registry.Category_Registry()
	BarangModules := barang_registry.Barang_Registry()
	TrxModules := trx_registry.Trx_Registry()

	auth_router.AuthRouter(app, AuthModules.Auth_Controller)
	toko_router.TokoRouter(app, TokoModules.TokoController)
	user_router.User_Register(app, UserModules.UserController)
	gudang_router.GudangRouter(app, GudangModules.GudangController)
	category_router.CategoryRouter(app, CategoryModules.CategoryController)
	barang_router.Barang_Router(app, BarangModules.BarangController)
	trx_router.Trx_Router(app, TrxModules.TrxController)

	app.Run(port_config.PORT)
}
