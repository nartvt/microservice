package orm

import (
	"health-service/app/domain/entities"
	"health-service/app/infra/db"
)

type IUser interface {
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByPhone(phoneNumber string) (*entities.User, error)
}
type user struct{}

var User IUser

func init() {
	User = user{}
}

func (u user) GetUserByEmail(email string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.Postgres.Model(&entities.User{}).
		Where("email = ?", email).
		Limit(1).
		Find(resp).Error
	return resp, err
}

func (u user) GetUserByPhone(phoneNumber string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.Postgres.Model(&entities.User{}).
		Where("phone_number = ?", phoneNumber).
		Limit(1).
		Find(resp).Error
	return resp, err
}
