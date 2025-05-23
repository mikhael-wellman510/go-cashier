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
	ProductController interface {
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	productController struct {
		productService usecases.ProductService
	}
)

func NewProductController(productService usecases.ProductService) ProductController {

	return &productController{
		productService: productService,
	}
}

func (pc *productController) Create(ctx *gin.Context) {

	var productReq entities.ProductRequest

	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}
	log.Println("Cek : ", productReq)
	res, err := pc.productService.CreateProduct(productReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Succes", res))

}

func (pc *productController) Update(ctx *gin.Context) {
	var productReq entities.ProductRequest

	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	pc.productService.UpdateProduct(productReq)
}
