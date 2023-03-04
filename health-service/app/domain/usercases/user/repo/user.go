package repo

import (
	"time"

	"gorm.io/gorm"

	"health-service/app/domain/usercases/user/orm"
	"health-service/app/uerror"
)

type UserRepo struct {
	Id          int
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	FullName    string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
type IUserRepo interface {
	GetUserByEmail(email string) (UserRepo, error)
	GetSectionsByPhone(phoneNumber string) (UserRepo, error)
}

type user struct{}

var User IUserRepo

func init() {
	User = user{}
}

func (u user) GetUserByEmail(email string) (UserRepo, error) {
	userOrm, err := orm.User.GetUserByEmail(email)
	if err != nil && err == gorm.ErrRecordNotFound {
		return UserRepo{}, nil
	}
	if err != nil {
		return UserRepo{}, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return UserRepo{}, nil
	}
	return UserRepo{
		Id:          userOrm.Id,
		UserName:    userOrm.UserName,
		FullName:    userOrm.FullName,
		Email:       userOrm.Email,
		PhoneNumber: userOrm.PhoneNumber,
		CreatedAt:   userOrm.CreatedAt,
		UpdatedAt:   userOrm.UpdatedAt,
	}, nil
}

func (u user) GetSectionsByPhone(phoneNumber string) (UserRepo, error) {
	userOrm, err := orm.User.GetUserByPhone(phoneNumber)
	if err != nil && err == gorm.ErrRecordNotFound {
		return UserRepo{}, nil
	}
	if err != nil {
		return UserRepo{}, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return UserRepo{}, nil
	}
	return UserRepo{
		Id:          userOrm.Id,
		UserName:    userOrm.UserName,
		FullName:    userOrm.FullName,
		Email:       userOrm.Email,
		PhoneNumber: userOrm.PhoneNumber,
		CreatedAt:   userOrm.CreatedAt,
		UpdatedAt:   userOrm.UpdatedAt,
	}, nil
}
