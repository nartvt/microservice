package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"health-service/app/domain/usercases/record/exercise/repo"
	"health-service/app/transport"
	"health-service/app/transport/record/exercise/request"
	"health-service/app/transport/record/exercise/response"
)

type ExerciseHandler struct {
	ExerciseDomain repo.IUseExerciseRepo
}

func (s ExerciseHandler) GetExerciseByUserId(ctx *fiber.Ctx) error {
	param := &request.UserExerciseParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}

	log.Printf("exercise-records %d\n", param.UserId)
	bodyRecords, err := s.ExerciseDomain.GetUserExerciseRepoRepoByUserId(param.UserId, param.Limit, param.Offset)
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
	nextUrl := fmt.Sprintf("%s?limit=%s&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
	resp.Pagination = &transport.Pagination{
		NextUrl: nextUrl,
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
