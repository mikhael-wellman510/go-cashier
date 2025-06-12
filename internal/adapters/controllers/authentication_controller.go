package controllers

import (
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthController interface {
		RegisterUserController(ctx *gin.Context)
		LoginController(ctx *gin.Context)
	}

	authController struct {
		authService usecases.AuthService
	}
)

func NewAuthController(authService usecases.AuthService) AuthController {

	return &authController{
		authService: authService,
	}
}
func (ac *authController) LoginController(ctx *gin.Context) {
	req := entities.LoginRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := ac.authService.Login(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("login succes", res))
}

func (ac *authController) RegisterUserController(ctx *gin.Context) {

	req := entities.UserRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := ac.authService.RegisterUser(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success Register", res))

}
