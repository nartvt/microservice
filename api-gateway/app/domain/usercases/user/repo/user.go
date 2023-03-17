package repo

import (
	"time"

	"gorm.io/gorm"

	"api-gateway/app/domain/entities"
	"api-gateway/app/domain/usercases/user/orm"
	"api-gateway/app/uerror"
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
	GetUserByEmail(email string) (*UserRepo, error)
	GetUserByUserName(userName string) (*UserRepo, error)
	GetUserById(id int) (*UserRepo, error)
	GetSectionsByPhone(phoneNumber string) (*UserRepo, error)
}

type user struct{}

var User IUserRepo

func init() {
	User = user{}
}

func (u user) GetUserById(id int) (*UserRepo, error) {
	userOrm, err := orm.User.GetUserById(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return nil, nil
	}
	return u.Bind(userOrm), nil
}

func (u user) GetUserByEmail(email string) (*UserRepo, error) {
	userOrm, err := orm.User.GetUserByEmail(email)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return nil, nil
	}
	return u.Bind(userOrm), nil
}

func (u user) GetUserByUserName(userName string) (*UserRepo, error) {
	userOrm, err := orm.User.GetUserByUserName(userName)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return nil, nil
	}
	return u.Bind(userOrm), nil
}

func (u user) GetSectionsByPhone(phoneNumber string) (*UserRepo, error) {
	userOrm, err := orm.User.GetUserByPhone(phoneNumber)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return nil, nil
	}
	return u.Bind(userOrm), nil
}

func (user) Bind(userEntity *entities.User) *UserRepo {
	if userEntity == nil {
		return nil
	}
	return &UserRepo{
		Id:          userEntity.Id,
		UserName:    userEntity.UserName,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		PhoneNumber: userEntity.PhoneNumber,
		CreatedAt:   userEntity.CreatedAt,
		UpdatedAt:   userEntity.UpdatedAt,
	}
}
