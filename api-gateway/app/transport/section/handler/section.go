package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api-gateway/app/domain/usercases/section/repo"
	"api-gateway/app/transport"
	"api-gateway/app/transport/section/request"
	"api-gateway/app/transport/section/response"
	"api-gateway/app/uerror"
)

type SectionHandler struct {
	SectionDomain repo.INewsfeedSectionRepo
}

func (s SectionHandler) GetSections(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("user invalid")})
	}
	param := &request.SectionRequestParam{}
	param.Bind(ctx)
	sections, err := s.SectionDomain.GetSections(param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}
	if len(sections) <= 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(uerror.NotFoundError(fmt.Errorf("not found %s", "section"), "section not found"))
	}
	return ctx.Status(http.StatusOK).JSON(transport.Response{
		Data: response.NewSectionResponses(sections),
	})
}

func (s SectionHandler) CreateSection(ctx *fiber.Ctx) error {
	param := &request.SectionRequestParam{}
	param.Bind(ctx)
	sections, err := s.SectionDomain.GetSections(param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}
	if len(sections) <= 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(uerror.NotFoundError(fmt.Errorf("not found %s", "section"), "section not found"))
	}
	return ctx.Status(http.StatusOK).JSON(transport.Response{
		Data: response.NewSectionResponses(sections),
	})
}
