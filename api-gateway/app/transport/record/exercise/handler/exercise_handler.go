package handler

import (
	"fmt"
	"log"
	"net/http"

	"api-gateway/app/domain/usercases/record/exercise/repo"
	uRepo "api-gateway/app/domain/usercases/user/repo"
	"api-gateway/app/transport"
	"api-gateway/app/transport/record/exercise/request"
	"api-gateway/app/transport/record/exercise/response"
)

type ExerciseHandler struct {
	ExerciseDomain repo.IUseExerciseRepo
}

func (s ExerciseHandler) GetExerciseByUserId(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("user invalid")})
	}
	userRepo := user.(uRepo.UserRepo)
	param := &request.UserExerciseParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}

	log.Printf("exercise-records %d\n", param.UserId)
	bodyRecords, err := s.ExerciseDomain.GetUserExerciseRepoRepoByUserId(userRepo.Id, param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}

	resp := transport.Response{
		Data: response.NewExeeciseResponses(bodyRecords),
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
