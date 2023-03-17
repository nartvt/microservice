package response

import (
	"api-gateway/app/domain/usercases/record/exercise/repo"
	"api-gateway/app/util"
)

func NewExeeciseResponse(diaryRepo repo.UserExerciseRepo) UserExerciesResponse {
	return UserExerciesResponse{
		Id:             diaryRepo.Id,
		UserId:         diaryRepo.UserId,
		AtTime:         diaryRepo.AtTime,
		Description:    diaryRepo.Description,
		CaloriesBurned: diaryRepo.CaloriesBurned,
		CreatedAt:      util.FormatDateTime(*diaryRepo.CreatedAt),
		UpdatedAt:      util.FormatDateTime(*diaryRepo.UpdatedAt),
	}
}
func NewExeeciseResponses(diaryRepos []repo.UserExerciseRepo) []UserExerciesResponse {
	if len(diaryRepos) <= 0 {
		return []UserExerciesResponse{}
	}
	resp := make([]UserExerciesResponse, len(diaryRepos))
	for i := range diaryRepos {
		resp[i] = NewExeeciseResponse(diaryRepos[i])
	}
	return resp
}
