package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"health-service/app/domain/usercases/section/repo"
	"health-service/app/transport"
	"health-service/app/transport/section/request"
	"health-service/app/transport/section/response"
	"health-service/app/uerror"
)

type SectionHandler struct {
	SectionDomain repo.INewsfeedSectionRepo
}

func (s SectionHandler) GetSections(ctx *fiber.Ctx) error {
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
