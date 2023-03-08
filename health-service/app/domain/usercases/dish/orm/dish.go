package orm

import (
	"gorm.io/gorm"

	"health-service/app/domain/entities"
	"health-service/app/domain/usercases/common"
	"health-service/app/infra/db"
)

type IDish interface {
	CreateDishTx(newDish *entities.Dish, tx *gorm.DB) error
	GetDishBySectionId(sectionId int, limit int, offset int) ([]entities.Dish, int, error)
}
type dish struct{}

var Dish IDish

func init() {
	Dish = dish{}
}
func (d dish) CreateDishTx(newDish *entities.Dish, tx *gorm.DB) error {
	return tx.Save(newDish).Error
}

func (d dish) GetDishBySectionId(sectionId int, limit int, offset int) ([]entities.Dish, int, error) {
	var resp []entities.Dish
	total := int64(0)
	err := db.DB().Model(&entities.Dish{}).
		InnerJoins("JOIN section_dishes ON dishes.id = section_dishes.dish_id").
		InnerJoins("JOIN newsfeed_sections ON section_dishes.section_id = newsfeed_sections.id").
		Where("newsfeed_sections.active = TRUE").
		Where("dishes.active = TRUE").
		Where("type = ?", common.NewsfeedSectionTypeTopPage).
		Where("newsfeed_sections.id = ?", sectionId).
		Limit(limit).
		Offset(offset).
		Order("id DESC").
		Count(&total).
		Find(&resp).Error
	return resp, int(total), err
}
