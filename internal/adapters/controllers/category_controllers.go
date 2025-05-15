package controllers

import (
	"fmt"
	"log"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/internal/utils"
	"mikhael-project-go/pkg/constants"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	CategoryController interface {
		CreateCategory(ctx *gin.Context)
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

	file, err := ctx.FormFile("image")

	log.Println("isi file : ", file)
	log.Println("isi req : ", categoryReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	fileName := fmt.Sprintf("uploads/%s%d%s", file.Filename, time.Now().Unix(), filepath.Ext(file.Filename))
	log.Println("Hasil fileName : ", fileName)
	if err := ctx.SaveUploadedFile(file, fileName); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	categoryReq.ImageUrl = constants.Server + fileName
	res, err := cc.categoryService.CreateCategory(categoryReq)

	log.Println("Hasil req : ", categoryReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}
