package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"health-service/app/domain/usercases/record/diary/repo"
	"health-service/app/transport"
	"health-service/app/transport/record/diary/request"
	"health-service/app/transport/record/diary/response"
)

type DiaryHandler struct {
	DiaryDomain repo.IUseDiaryRepo
}

func (s DiaryHandler) GetDiariesByUserId(ctx *fiber.Ctx) error {
	param := &request.DiaryRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}
	log.Printf("diary-records %d\n", param.UserId)
	bodyRecords, err := s.DiaryDomain.GetUserDiaryRepoByUserId(param.UserId, param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}

	resp := transport.Response{
		Data: response.NewDiaryResponses(bodyRecords),
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
