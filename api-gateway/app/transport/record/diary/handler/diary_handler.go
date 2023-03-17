package handler

import (
	"fmt"
	"log"
	"net/http"

	"api-gateway/app/domain/usercases/record/diary/repo"
	uRepo "api-gateway/app/domain/usercases/user/repo"
	"api-gateway/app/transport"
	"api-gateway/app/transport/record/diary/request"
	"api-gateway/app/transport/record/diary/response"
)

type DiaryHandler struct {
	DiaryDomain repo.IUseDiaryRepo
}

func (s DiaryHandler) GetDiariesByUserId(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("user invalid")})
	}
	userRepo := user.(uRepo.UserRepo)
	param := &request.DiaryRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}
	log.Printf("diary-records %d\n", param.UserId)
	bodyRecords, err := s.DiaryDomain.GetUserDiaryRepoByUserId(userRepo.Id, param.Limit, param.Offset)
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
	nextUrl := fmt.Sprintf("%s?limit=%d&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
	resp.Pagination = &transport.Pagination{
		NextUrl: nextUrl,
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
