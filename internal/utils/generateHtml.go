package utils

import (
	"fmt"
	"mikhael-project-go/internal/entities"
)

func GenerateHtml(data []entities.ProductResponse) string {

	html := `<table border="1" cellpadding="5" cellspacing="0" style="border-collapse: collapse;">`
	html += `<tr><th>ID</th><th>Stock</th><th>Barcode</th><th>Product Name</th><th>Description</th><th>Price</th><th>Category Name</th><th>Category Description</th><th>Image URL</th><th>Store Name</th><th>Address</th><th>Owner Name</th></tr>`

	for _, product := range data {
		html += fmt.Sprintf("<tr><td>%s</td><td>%d</td><td>%s</td><td>%s</td><td>%s</td><td>%f</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", product.Id, product.Stock, product.Barcode, product.ProductName, product.Description, product.Price.InexactFloat64(), product.CategoryResponse.CategoryName, product.CategoryResponse.Description, product.CategoryResponse.ImageUrl, product.StoreResponse.StoreName, product.StoreResponse.Address, product.StoreResponse.OwnerName)

	}

	html += "</table>"

	return html
}
