package orm

import (
	"api-gateway/app/domain/entities"
	"api-gateway/app/infra/db"
)

type IUser interface {
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByPhone(phoneNumber string) (*entities.User, error)
	GetUserByUserName(userName string) (*entities.User, error)
}
type user struct{}

var User IUser

func init() {
	User = user{}
}

func (u user) GetUserById(id int) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("id = ?", id).
		Limit(1).
		Find(resp).Error
	return resp, err
}

func (u user) GetUserByEmail(email string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("email = ?", email).
		Limit(1).
		Find(resp).Error
	return resp, err
}
func (u user) GetUserByUserName(userName string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("user_name = ?", userName).
		Limit(1).
		Find(resp).Error
	return resp, err
}
func (u user) GetUserByPhone(phoneNumber string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("phone_number = ?", phoneNumber).
		Limit(1).
		Find(resp).Error
	return resp, err
}
