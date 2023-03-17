package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api-gateway/app/domain/usercases/product/repo"
	"api-gateway/app/transport"
	"api-gateway/app/transport/product/request"
	"api-gateway/app/transport/product/response"
	"api-gateway/app/uerror"
)

type ProductHandler struct {
	ProductDomain repo.IProductRepo
}

func (s ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	input := &request.ProductRequestBody{}
	if err := input.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	product, err := s.ProductDomain.CreateProductTx(*input)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(transport.Response{
		Data: response.NewProductResponse(product),
	})
}
func (s ProductHandler) GetProductBySectionId(ctx *fiber.Ctx) error {
	param := &request.ProductRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}
	products, total, err := s.ProductDomain.GetProductBySectionId(param.SectionId, param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}

	if len(products) <= 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(uerror.NotFoundError(fmt.Errorf("not found %s", "section"), "section not found"))
	}

	resp := transport.Response{
		Data: response.NewProductResponses(products),
	}
	if len(products) >= param.Limit {
		nextUrl := fmt.Sprintf("%s?limit=%d&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
		resp.Pagination = &transport.Pagination{
			NextUrl: nextUrl,
			Total:   total,
		}
		return ctx.Status(http.StatusOK).JSON(resp)
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
