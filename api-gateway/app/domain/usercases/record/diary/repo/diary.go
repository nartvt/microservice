package repo

import (
	"time"

	"api-gateway/app/domain/entities"
	"api-gateway/app/domain/usercases/record/diary/orm"
	"api-gateway/app/uerror"
)

type UserDiaryRepo struct {
	Id          int
	UserId      int
	AtTime      int
	Description string
	Calories    int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
type IUseDiaryRepo interface {
	GetUserDiaryRepoByUserId(userId int, limit int, offset int) ([]UserDiaryRepo, error)
}
type userBodyRecordRepo struct{}

func NewUserBodyRecordRepoRepo() *userBodyRecordRepo {
	return &userBodyRecordRepo{}
}

func (u userBodyRecordRepo) GetUserDiaryRepoByUserId(userId int, limit int, offset int) ([]UserDiaryRepo, error) {
	userDiaries, err := orm.UserDiary.GetUserDiaryByUserId(userId, limit, offset)
	if err != nil {
		return []UserDiaryRepo{}, uerror.InternalError(err, err.Error())
	}
	return u.populateUserDiaries(userDiaries), nil
}

func (u userBodyRecordRepo) populateUserDiaries(userDiaries []entities.UserDiary) []UserDiaryRepo {
	resp := make([]UserDiaryRepo, len(userDiaries))
	for i := range userDiaries {
		resp[i] = u.populateUserDiary(userDiaries[i])
	}
	return resp
}
func (userBodyRecordRepo) populateUserDiary(userDiaryEntity entities.UserDiary) UserDiaryRepo {
	resp := UserDiaryRepo{
		Id:          userDiaryEntity.Id,
		UserId:      userDiaryEntity.UserId,
		AtTime:      userDiaryEntity.AtTime,
		Description: userDiaryEntity.Description,
		Calories:    userDiaryEntity.Calories,
	}
	if userDiaryEntity.CreatedAt != nil {
		resp.CreatedAt = userDiaryEntity.CreatedAt
	}
	if userDiaryEntity.UpdatedAt != nil {
		resp.UpdatedAt = userDiaryEntity.UpdatedAt
	}
	return resp
}
