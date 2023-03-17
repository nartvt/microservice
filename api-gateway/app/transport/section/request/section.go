package request

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/app/transport"
)

type SectionRequestParam struct {
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

func (input *SectionRequestParam) Bind(c *fiber.Ctx) error {
	limit := c.QueryInt(transport.ParamLimit, transport.DefaultLimit)
	page := c.QueryInt(transport.ParamPage, transport.DefgaultPage)
	offset := limit * (page - 1)
	input.Page = page
	input.Limit = limit
	input.Offset = offset
	return nil
}
