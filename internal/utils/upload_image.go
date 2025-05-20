package utils

import (
	"fmt"
	"log"
	"path"

	"mikhael-project-go/pkg/constants"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// Todo -> form field untuk isi req di form
func UploadImage(ctx *gin.Context, formField string, uploadDir string) (string, error) {

	// Todo -> dapatkan file nya
	file, err := ctx.FormFile(formField)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return "", err
	}

	// Todo -> cek apakah jpg / png / svg
	ext := filepath.Ext(file.Filename)
	log.Println("ext apa ? ", ext)
	fileName := fmt.Sprintf("%s%d%s", file.Filename, time.Now().Unix(), ext)
	log.Println("hasil file name :", fileName)
	savePath := path.Join(uploadDir, fileName)

	log.Println("hasil save :  ", savePath)

	// Todo -> Save ke foto ke file system
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		// ctx.JSON(http.StatusInternalServerError, BuildResponseFailed(err.Error()))
		return "", err
	}

	return constants.Server + savePath, nil
}
