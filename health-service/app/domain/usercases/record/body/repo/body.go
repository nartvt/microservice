package repo

import (
	"time"

	"health-service/app/domain/entities"
	"health-service/app/domain/usercases/record/body/orm"
	"health-service/app/uerror"
)

type UserBodyRecordRepo struct {
	Id         int
	UserId     int
	Weight     float32
	Height     int
	Percentage float32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type IUserBodyRecordRepo interface {
	GetUserBodyRecordRepoByUserId(userId int, limit int, offset int) ([]UserBodyRecordRepo, error)
}
type userBodyRecordRepo struct{}

func NewUserBodyRecordRepo() *userBodyRecordRepo {
	return &userBodyRecordRepo{}
}

func (u userBodyRecordRepo) GetUserBodyRecordRepoByUserId(userId int, limit int, offset int) ([]UserBodyRecordRepo, error) {
	bodyEntities, err := orm.UserBodyRecord.GetBodyRecordByUserId(userId, limit, offset)
	if err != nil {
		return []UserBodyRecordRepo{}, uerror.InternalError(err, err.Error())
	}
	return u.populateUserBodyRecords(bodyEntities), nil
}

func (u userBodyRecordRepo) populateUserBodyRecords(userBodyRecordEntities []entities.UserBodyRecord) []UserBodyRecordRepo {
	resp := make([]UserBodyRecordRepo, len(userBodyRecordEntities))
	for i := range userBodyRecordEntities {
		resp[i] = u.populateUserBodyRecord(userBodyRecordEntities[i])
	}
	return resp
}
func (userBodyRecordRepo) populateUserBodyRecord(userBodyRecordEntity entities.UserBodyRecord) UserBodyRecordRepo {
	resp := UserBodyRecordRepo{
		Id:         userBodyRecordEntity.Id,
		UserId:     userBodyRecordEntity.UserId,
		Weight:     userBodyRecordEntity.Weight,
		Height:     userBodyRecordEntity.Height,
		Percentage: userBodyRecordEntity.Percentage,
	}
	if userBodyRecordEntity.CreatedAt != nil {
		resp.CreatedAt = userBodyRecordEntity.CreatedAt
	}
	if userBodyRecordEntity.UpdatedAt != nil {
		resp.UpdatedAt = userBodyRecordEntity.UpdatedAt
	}
	return resp
}
