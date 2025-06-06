package controllers

import (
	"log"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	StoreController interface {
		CreateStore(ctx *gin.Context)
		FindStoreById(ctx *gin.Context)
		UpdateStore(ctx *gin.Context)
		DeletedStore(ctx *gin.Context)
		GetStoreByPaggingAndFilter(ctx *gin.Context)
	}

	storeController struct {
		storeService usecases.StoreService
	}
)

func NewStoreController(storeService usecases.StoreService) StoreController {

	return &storeController{
		storeService: storeService,
	}
}

func (sc *storeController) CreateStore(ctx *gin.Context) {
	var post entities.StoreRequest

	if err := ctx.ShouldBind(&post); err != nil {

		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	log.Println("Controller : ", post)

	res, err := sc.storeService.CreateStore(&post)

	log.Println("Hasil controller res : ", res)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (sc *storeController) FindStoreById(ctx *gin.Context) {
	params := ctx.Param("id")

	log.Println("params : ", params)
	if params == "" {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed("Id tidak di isi"))
		return
	}
	res, err := sc.storeService.FindStoreById(params)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.BuildResponseFailed("Id tidak di temukan"))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))

}

func (sc *storeController) UpdateStore(ctx *gin.Context) {
	var updates entities.StoreRequest

	if err := ctx.ShouldBind(&updates); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := sc.storeService.UpdatedStore(&updates)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Updated Succes", res))

}

func (sc *storeController) DeletedStore(ctx *gin.Context) {
	params := ctx.Param("id")

	res, err := sc.storeService.DeletedStore(params)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success Deleted id : "+params, res))

}

func (sc *storeController) GetStoreByPaggingAndFilter(ctx *gin.Context) {
	storeName := ctx.DefaultQuery("storeName", "")
	ownerName := ctx.DefaultQuery("ownerName", "")

	paggination := utils.GetPagination(ctx)

	res, err := sc.storeService.FilterAndPagginStore(paggination.Page, paggination.Limit, storeName, ownerName)
	// Todo getQuery

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Succes ", res))
}
