package controllers

import (
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/usecases"
	"mikhael-project-go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	PaymentMethodControler interface {
		CreatePaymentMethodController(ctx *gin.Context)
		FindAllPaymentMethodController(ctx *gin.Context)
		DeletedPaymentMethodController(ctx *gin.Context)
	}

	paymentMethodController struct {
		paymentMethodService usecases.PaymentMethodService
	}
)

func NewPaymentMethodController(paymentMethodService usecases.PaymentMethodService) PaymentMethodControler {

	return &paymentMethodController{
		paymentMethodService: paymentMethodService,
	}
}

func (pmc *paymentMethodController) CreatePaymentMethodController(ctx *gin.Context) {

	var paymentMethodReq entities.PaymentMethodRequest

	if err := ctx.ShouldBind(&paymentMethodReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := pmc.paymentMethodService.CreatePaymentMethod(&paymentMethodReq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Succes", res))
}

func (pmc *paymentMethodController) FindAllPaymentMethodController(ctx *gin.Context) {

	res, err := pmc.paymentMethodService.FindAllPaymentMethod()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (pmc *paymentMethodController) DeletedPaymentMethodController(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := pmc.paymentMethodService.DeletedPaymentMethod(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success Deleted id : "+id, res))
}
