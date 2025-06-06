package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	ProductController interface {
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		FindById(ctx *gin.Context)
		PaggingProduct(ctx *gin.Context)
		ExportProductToCsv(ctx *gin.Context)
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
	res, err := pc.productService.CreateProduct(&productReq)

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

	log.Println("Error product controler  : ", productReq)

	res, err := pc.productService.UpdateProduct(&productReq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Succes ", res))
}

func (pc *productController) FindById(ctx *gin.Context) {
	params := ctx.Param("id")

	res, err := pc.productService.FindProductById(params)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Succes", res))
}

func (pc *productController) PaggingProduct(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")

	// Todo -> ini untuk handle jika page / limit tidak valid
	page := utils.GetPagination(ctx)

	res, err := pc.productService.FilterAndPaggingProduct(page.Page, page.Limit, search)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Succes", res))
}

func (pc *productController) ExportProductToCsv(ctx *gin.Context) {
	date := ctx.Query("date")

	res, err := pc.productService.ExportProductToCsv(date)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=products_%s.csv", date))
	ctx.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(ctx.Writer)
	defer writer.Flush()

	writer.Write([]string{"ID", "Stock", "Barcode", "ProductName", "Description", "Price", "CategoryName", "Description", "ImageUrl", "StoreName", "Address", "OwnerName"})

	for _, p := range res {
		record := []string{
			p.Id,
			strconv.Itoa(p.Stock),
			p.Barcode,
			p.ProductName,
			p.Description,
			p.Price.String(),
			p.CategoryResponse.CategoryName,
			p.CategoryResponse.Description,
			p.CategoryResponse.ImageUrl,
			p.StoreResponse.StoreName,
			p.StoreResponse.Address,
			p.StoreResponse.OwnerName,
		}
		writer.Write(record)
	}

}
