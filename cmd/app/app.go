package app

import (
	"fmt"
	"log"
	"mikhael-project-go/config"
	"mikhael-project-go/internal/adapters/controllers"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/service"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/migrations"
	"mikhael-project-go/pkg/constants"
	"mikhael-project-go/pkg/drivers/sql"
	"mikhael-project-go/seeders"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Db     *gorm.DB
	Router *gin.Engine
}

func (app *App) ConnectDb() {
	db := sql.OpenDbConnection()
	app.Db = db
	err := migrations.Migrate(db)

	if err != nil {
		log.Fatal("Migrations Failed : -> ", err)
	}

}

func (app *App) Routes() {

	router := gin.Default()
	router.Static("/uploads", "./uploads")

	baseUrl := fmt.Sprintf("%s/v%d", constants.ApiPrevix, constants.ApiVersion)

	authRepo := repositories.NewUserRepository(app.Db)
	authUseCase := usecases.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authUseCase)

	authRoutes := router.Group(fmt.Sprintf("%s/auth", baseUrl))
	authRoutes.POST("/register", authController.RegisterUserController)
	authRoutes.POST("/login", authController.LoginController)

	storeRepo := repositories.NewStoreRepository(app.Db)
	storeUseCase := usecases.NewStoreService(storeRepo)
	storeController := controllers.NewStoreController(storeUseCase)

	storeRoutes := router.Group(fmt.Sprintf("%s/store", baseUrl))
	storeRoutes.POST("/create", storeController.CreateStore)
	storeRoutes.GET("/find/:id", storeController.FindStoreById)
	storeRoutes.GET("/find", storeController.FindStoreById)
	storeRoutes.PUT("/update", storeController.UpdateStore)
	storeRoutes.DELETE("/deleted/:id", storeController.DeletedStore)
	storeRoutes.GET("/searchAndFilterStore", storeController.GetStoreByPaggingAndFilter)

	categoryRepo := repositories.NewCategoriesRepository(app.Db)
	categoryUseCase := usecases.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryUseCase)

	categoryRoutes := router.Group(fmt.Sprintf("%s/category", baseUrl))
	categoryRoutes.POST("/create", categoryController.CreateCategory)
	categoryRoutes.PUT("/updated", categoryController.UpdateCategory)

	productRepo := repositories.NewProductRepository(app.Db)
	productUseCase := usecases.NewProductService(productRepo, storeUseCase, categoryUseCase)
	productController := controllers.NewProductController(productUseCase)

	productRoutes := router.Group(fmt.Sprintf("%s/product", baseUrl))
	productRoutes.POST("/create", productController.Create)
	productRoutes.PUT("/updated", productController.Update)
	productRoutes.GET("/find/:id", productController.FindById)
	productRoutes.GET("/search", productController.PaggingProduct)
	productRoutes.GET("/exportCsv", productController.ExportProductToCsv)

	// Scheduler service
	schedulerUseCase := usecases.NewSchedulerService(productUseCase)

	// Seeders
	seed := seeders.NewSeeders(app.Db)
	router.GET("/seeds", seed.GenerateSeeders)

	// Run Cron
	cronJob := service.NewCronJob(schedulerUseCase)
	cronJob.StartSchedulerSendEmail()
	app.Router = router

}

// Running port nya
func (app *App) Run() {
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	app.Router.Run(port)
}
