package usecases

import (
	"fmt"

	"mikhael-project-go/internal/utils"
)

type (
	SchedulerService interface {
		SendEmailProduct()
	}

	schedulerService struct {
		ProductService ProductService
	}
)

func NewSchedulerService(productService ProductService) SchedulerService {

	return &schedulerService{
		ProductService: productService,
	}
}

func (ss *schedulerService) SendEmailProduct() {
	// get data dari product service
	res, err := ss.ProductService.ExportProductToCsv("2025-05-28")

	htmlTable := utils.GenerateHtml(res)
	if err != nil {
		fmt.Println("Error fetch data !")
		return
	}
	utils.SendEmail("mikhaelwellman423@gmail.com", "List Product", htmlTable)

}
