package handler

import (
	"fmt"
	"net/http"

	uRepo "api-gateway/app/domain/usercases/user/repo"

	"api-gateway/app/domain/usercases/record/body/repo"
	"api-gateway/app/transport"
	"api-gateway/app/transport/record/body/request"
	"api-gateway/app/transport/record/body/response"
)

type BodyRecordHandler struct {
	BodyRecordDomain repo.IUserBodyRecordRepo
}

func (s BodyRecordHandler) GetBodyRecordsByUserId(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("user invalid")})
	}
	userRepo := user.(uRepo.UserRepo)
	param := &request.BodyRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	bodyRecords, err := s.BodyRecordDomain.GetUserBodyRecordRepoByUserId(userRepo.Id, param.Limit, param.Offset)
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
	nextUrl := fmt.Sprintf("%s?limit=%d&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
	resp.Pagination = &transport.Pagination{
		NextUrl: nextUrl,
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
