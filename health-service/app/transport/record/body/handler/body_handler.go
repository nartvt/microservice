package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"health-service/app/domain/usercases/record/body/repo"
	"health-service/app/transport"
	"health-service/app/transport/record/body/request"
	"health-service/app/transport/record/body/response"
)

type BodyRecordHandler struct {
	BodyRecordDomain repo.IUserBodyRecordRepo
}

func (s BodyRecordHandler) GetBodyRecordsByUserId(ctx *fiber.Ctx) error {
	param := &request.BodyRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}
	log.Printf("body-records %d\n", param.UserId)
	bodyRecords, err := s.BodyRecordDomain.GetUserBodyRecordRepoByUserId(param.UserId, param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}

	resp := transport.Response{
		Data: response.NewBodyRecordResponses(bodyRecords),
	}
	if len(bodyRecords) < param.Limit {
		return ctx.Status(http.StatusOK).JSON(resp)
	}
	nextUrl := fmt.Sprintf("%s?limit=%s&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
	resp.Pagination = &transport.Pagination{
		NextUrl: nextUrl,
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
