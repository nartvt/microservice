package orm

import (
	"health-service/app/domain/entities"
	"health-service/app/infra/db"
)

type INewsfeedSection interface {
	GetSections(limit int, offset int) ([]entities.NewsfeedSection, error)
	GetSectionsById(id int) (*entities.NewsfeedSection, error)
}
type newsfeedSection struct{}

var NewsfeedSection INewsfeedSection

func init() {
	NewsfeedSection = newsfeedSection{}
}

func (n newsfeedSection) GetSections(limit int, offset int) ([]entities.NewsfeedSection, error) {
	resp := []entities.NewsfeedSection{}
	err := db.DB().Model(&entities.NewsfeedSection{}).
		Where("active = TRUE").
		Limit(limit).
		Offset(offset).
		Order("id DESC").
		Find(&resp).Error
	return resp, err
}

func (n newsfeedSection) GetSectionsById(id int) (*entities.NewsfeedSection, error) {
	resp := &entities.NewsfeedSection{}
	err := db.DB().Model(&entities.NewsfeedSection{}).
		Where("active = TRUE").
		Where("id = ?", id).
		Find(resp).
		Error
	return resp, err
}
