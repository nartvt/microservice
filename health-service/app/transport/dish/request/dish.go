package request

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"health-service/app/transport"
	"health-service/app/uerror"
)

type DishRequestParam struct {
	Limit     int `json:"limit"`
	Page      int `json:"page"`
	Offset    int `json:"offset"`
	SectionId int `json:"section_id"`
}

type DishRequestBody struct {
	Name   string
	Active bool
	Image  string
}

func (input *DishRequestBody) Bind(ctx *fiber.Ctx) error {
	err := json.Unmarshal(ctx.Body(), input)
	if err != nil {
		return uerror.BadRequestError(err, err.Error())
	}
	return nil
}

func (input *DishRequestParam) Bind(c *fiber.Ctx) error {
	limit := c.QueryInt(transport.ParamLimit, transport.DefaultLimit)
	page := c.QueryInt(transport.ParamPage, transport.DefgaultPage)
	sectionId, err := c.ParamsInt(transport.ParamSectionId, 0)
	if err != nil {
		return uerror.BadRequestError(err, fmt.Sprintf("section invalid %s", "sectionId"))
	}
	if sectionId <= 0 {
		return uerror.BadRequestError(fmt.Errorf("section invalid %s", "sectionId"), "section invalid")
	}
	offset := limit * (page - 1)
	input.Page = page
	input.Limit = limit
	input.Offset = offset
	input.SectionId = sectionId
	return nil
}
