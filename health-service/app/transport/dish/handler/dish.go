package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"health-service/app/domain/usercases/dish/repo"
	"health-service/app/transport"
	"health-service/app/transport/dish/request"
	"health-service/app/transport/dish/response"
	"health-service/app/uerror"
)

type DishHandler struct {
	DishDomain repo.IDishRepo
}

func (s DishHandler) CreateDish(ctx *fiber.Ctx) error {
	input := &request.DishRequestBody{}
	if err := input.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	dish, err := s.DishDomain.CreateDishTx(*input)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(transport.Response{
		Data: response.NewDishResponse(dish),
	})
}
func (s DishHandler) GetDishBySectionId(ctx *fiber.Ctx) error {
	param := &request.DishRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}
	dishes, total, err := s.DishDomain.GetDishBySectionId(param.SectionId, param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}

	if len(dishes) <= 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(uerror.NotFoundError(fmt.Errorf("not found %s", "section"), "section not found"))
	}

	resp := transport.Response{
		Data: response.NewDishResponses(dishes),
	}
	if len(dishes) < param.Limit {
		return ctx.Status(http.StatusOK).JSON(resp)
	}
	nextUrl := fmt.Sprintf("%s?limit=%s&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
	resp.Pagination = &transport.Pagination{
		NextUrl: nextUrl,
		Total:   total,
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
