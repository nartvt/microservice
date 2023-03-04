package request

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"health-service/app/transport"
	"health-service/app/uerror"
)

type UserExerciseParam struct {
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
	UserId int `json:"section_id"`
}

func (input *UserExerciseParam) Bind(c *fiber.Ctx) error {
	limit := c.QueryInt(transport.ParamLimit, transport.DefaultLimit)
	page := c.QueryInt(transport.ParamPage, transport.DefgaultPage)
	userId, err := c.ParamsInt(transport.ParamUserId, 0)
	if err != nil {
		return uerror.BadRequestError(err, fmt.Sprintf("user invalid %s", "userId"))
	}
	if userId <= 0 {
		return uerror.BadRequestError(fmt.Errorf("user id invalid %s", "userId"), "user invalid")
	}
	offset := limit * (page - 1)
	input.Page = page
	input.Limit = limit
	input.Offset = offset
	input.UserId = userId
	return nil
}
