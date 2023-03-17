package response

import (
	"api-gateway/app/domain/usercases/record/diary/repo"
	"api-gateway/app/util"
)

func NewDiaryResponse(diaryRepo repo.UserDiaryRepo) UserDiaryResponse {
	return UserDiaryResponse{
		Id:          diaryRepo.Id,
		UserId:      diaryRepo.UserId,
		AtTime:      diaryRepo.AtTime,
		Description: diaryRepo.Description,
		Calories:    diaryRepo.Calories,
		CreatedAt:   util.FormatDateTime(*diaryRepo.CreatedAt),
		UpdatedAt:   util.FormatDateTime(*diaryRepo.UpdatedAt),
	}
}

func NewDiaryResponses(diaryRepos []repo.UserDiaryRepo) []UserDiaryResponse {
	if len(diaryRepos) <= 0 {
		return []UserDiaryResponse{}
	}
	resp := make([]UserDiaryResponse, len(diaryRepos))
	for i := range diaryRepos {
		resp[i] = NewDiaryResponse(diaryRepos[i])
	}
	return resp
}
