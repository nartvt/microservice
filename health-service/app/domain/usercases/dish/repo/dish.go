package repo

import (
	"time"

	"gorm.io/gorm"

	"health-service/app/domain/entities"
	"health-service/app/domain/usercases/common"
	"health-service/app/domain/usercases/dish/orm"
	"health-service/app/transport/dish/request"
	"health-service/app/uerror"
)

type DishRepo struct {
	Id        int
	Name      string
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}
type IDishRepo interface {
	CreateDishTx(newDish request.DishRequestBody) (DishRepo, error)
	GetDishBySectionId(sectionId int, limit int, offset int) ([]DishRepo, int, error)
}
type dishRepo struct{}

func NewDishRepo() *dishRepo {
	return &dishRepo{}
}

func (dishRepo) CreateDishTx(newDish request.DishRequestBody) (DishRepo, error) {
	tx := common.BeginTx()
	defer common.RecoveryTx(tx)

	now := time.Now()
	newDishEntity := entities.Dish{
		Name:      newDish.Name,
		Active:    true,
		Image:     newDish.Image,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := orm.Dish.CreateDishTx(&newDishEntity, tx)
	if err != nil {
		tx.Rollback()
		return DishRepo{}, uerror.InternalError(err, err.Error())
	}
	return DishRepo{
		Id:        newDishEntity.Id,
		Active:    newDishEntity.Active,
		Image:     newDishEntity.Image,
		CreatedAt: newDishEntity.CreatedAt,
		UpdatedAt: newDishEntity.UpdatedAt,
	}, nil
}

func (d dishRepo) GetDishBySectionId(sectionId int, limit int, offset int) ([]DishRepo, int, error) {
	dishes, total, err := orm.Dish.GetDishBySectionId(sectionId, limit, offset)
	if err != nil && err == gorm.ErrRecordNotFound {
		return []DishRepo{}, 0, nil
	}
	if err != nil {
		return []DishRepo{}, 0, uerror.InternalError(err, err.Error())
	}
	return d.populateDishes(dishes), total, nil
}

func (d dishRepo) populateDishes(dishes []entities.Dish) []DishRepo {
	if len(dishes) <= 0 {
		return []DishRepo{}
	}
	resp := make([]DishRepo, len(dishes))
	for i := range dishes {
		resp[i] = d.populateDish(dishes[i])
	}
	return resp
}

func (d dishRepo) populateDish(dish entities.Dish) DishRepo {
	return DishRepo{
		Id:        dish.Id,
		Name:      dish.Name,
		Image:     dish.Image,
		Active:    dish.Active,
		CreatedAt: dish.CreatedAt,
		UpdatedAt: dish.UpdatedAt,
	}
}
