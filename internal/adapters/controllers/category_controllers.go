package controllers

import (
	"log"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/internal/utils"
	"mikhael-project-go/pkg/constants"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	CategoryController interface {
		CreateCategory(ctx *gin.Context)
		UpdateCategory(ctx *gin.Context)
	}

	categoryController struct {
		categoryService usecases.CategoryService
	}
)

func NewCategoryController(categoryService usecases.CategoryService) CategoryController {

	return &categoryController{
		categoryService: categoryService,
	}
}

func (cc *categoryController) CreateCategory(ctx *gin.Context) {

	var categoryReq entities.CategoryRequest

	if err := ctx.ShouldBindWith(&categoryReq, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	// todo -> upload imahe
	// image -> ini nama file form nya
	// uploads -> nama folder nya mau di taro dmna !
	val, err := utils.UploadImage(ctx, "image", "uploads")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	categoryReq.ImageUrl = val
	res, err := cc.categoryService.CreateCategory(&categoryReq)

	if err != nil {
		log.Println("hasil : ", strings.TrimPrefix(categoryReq.ImageUrl, constants.Server))
		_ = os.Remove(strings.TrimPrefix(categoryReq.ImageUrl, constants.Server))
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (cc *categoryController) UpdateCategory(ctx *gin.Context) {
	var categoryReq entities.CategoryRequest

	if err := ctx.ShouldBindWith(&categoryReq, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}
	// todo -> upload gambar baru
	val, err := utils.UploadImage(ctx, "image", "uploads")
	categoryReq.ImageUrl = val

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := cc.categoryService.UpdateCategory(&categoryReq)

	if err != nil {
		// Handle untuk hapus foto jika gagal dalam transaksi
		_ = os.Remove(strings.TrimPrefix(val, constants.Server))
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Succes", res))

}
