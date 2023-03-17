package orm

import (
	"api-gateway/app/domain/entities"
	"api-gateway/app/infra/db"
)

type IUserExercise interface {
	GetUserExciseByUserId(userId int, limit int, offset int) ([]entities.UserExercise, error)
}
type userExercise struct{}

var UserExercise IUserExercise

func init() {
	UserExercise = userExercise{}
}
func (u userExercise) GetUserExciseByUserId(userId int, limit int, offset int) ([]entities.UserExercise, error) {
	if userId <= 0 {
		return []entities.UserExercise{}, nil
	}
	resp := []entities.UserExercise{}
	err := db.DB().Model(&entities.UserExercise{}).
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).Order("id DESC").
		Find(&resp).Error
	return resp, err
}
