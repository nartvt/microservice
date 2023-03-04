package repo

import (
	"time"

	"health-service/app/domain/entities"
	"health-service/app/domain/usercases/record/exercise/orm"
	"health-service/app/uerror"
)

type UserExerciseRepo struct {
	Id             int
	UserId         int
	AtTime         int
	Description    string
	CaloriesBurned int
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
type IUseExerciseRepo interface {
	GetUserExerciseRepoRepoByUserId(userId int, limit int, offset int) ([]UserExerciseRepo, error)
}
type userExerciseRepo struct{}

func NewUserExerciseRecordRepoRepo() *userExerciseRepo {
	return &userExerciseRepo{}
}

func (u userExerciseRepo) GetUserExerciseRepoRepoByUserId(userId int, limit int, offset int) ([]UserExerciseRepo, error) {
	userExercises, err := orm.UserExercise.GetUserExciseByUserId(userId, limit, offset)
	if err != nil {
		return []UserExerciseRepo{}, uerror.InternalError(err, err.Error())
	}
	return u.populateUserExercies(userExercises), nil
}

func (u userExerciseRepo) populateUserExercies(userExercises []entities.UserExercise) []UserExerciseRepo {
	resp := make([]UserExerciseRepo, len(userExercises))
	for i := range userExercises {
		resp[i] = u.populateUserExercie(userExercises[i])
	}
	return resp
}
func (userExerciseRepo) populateUserExercie(userDiaryEntity entities.UserExercise) UserExerciseRepo {
	resp := UserExerciseRepo{
		Id:             userDiaryEntity.Id,
		UserId:         userDiaryEntity.UserId,
		AtTime:         userDiaryEntity.AtTime,
		Description:    userDiaryEntity.Description,
		CaloriesBurned: userDiaryEntity.CaloriesBurned,
	}
	if userDiaryEntity.CreatedAt != nil {
		resp.CreatedAt = userDiaryEntity.CreatedAt
	}
	if userDiaryEntity.UpdatedAt != nil {
		resp.UpdatedAt = userDiaryEntity.UpdatedAt
	}
	return resp
}
