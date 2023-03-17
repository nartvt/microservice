package response

import (
	"api-gateway/app/domain/usercases/product/repo"
	"api-gateway/app/util"
)

func NewProductResponse(product repo.ProductRepo) ProductModelResponse {
	productResp := ProductModelResponse{
		Id:    product.Id,
		Name:  product.Name,
		Image: product.Image,
	}
	if product.CreatedAt != nil {
		productResp.CreatedAt = util.FormatDateTime(*product.CreatedAt)
	}
	if product.UpdatedAt != nil {
		productResp.UpdatedAt = util.FormatDateTime(*product.UpdatedAt)
	}
	return productResp
}

func NewProductResponses(products []repo.ProductRepo) []ProductModelResponse {
	if len(products) <= 0 {
		return []ProductModelResponse{}
	}
	productResp := make([]ProductModelResponse, len(products))
	for i := range products {
		productResp[i] = NewProductResponse(products[i])
	}
	return productResp
}
