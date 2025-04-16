package app

import (
	"fmt"
	"log"
	"mikhael-project-go/config"
	"mikhael-project-go/internal/adapters/controllers"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/migrations"
	"mikhael-project-go/pkg/constants"
	"mikhael-project-go/pkg/drivers/sql"

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

	baseUrl := fmt.Sprintf("%s/v%d", constants.ApiPrevix, constants.ApiVersion)
	log.Println("Base url : ", baseUrl)
	storeRepo := repositories.NewStoreRepository(app.Db)
	storeUseCase := usecases.NewStoreService(storeRepo)
	storeController := controllers.NewStoreController(storeUseCase)

	storeRoutes := router.Group(fmt.Sprintf("%s/product", baseUrl))
	storeRoutes.POST("/create", storeController.CreateStore)
	storeRoutes.GET("/find/:id", storeController.FindStoreById)
	storeRoutes.GET("/find", storeController.FindStoreById)
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"Message": "Crot",
	// 	})
	// })

	app.Router = router
}

// Running port nya
func (app *App) Run() {
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	app.Router.Run(port)
}
