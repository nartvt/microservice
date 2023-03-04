package repo

import (
	"time"

	"gorm.io/gorm"

	"health-service/app/domain/entities"
	"health-service/app/domain/usercases/section/orm"
	"health-service/app/uerror"
)

type NewsfeedSectionRepo struct {
	Id        int
	Name      string
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}
type INewsfeedSectionRepo interface {
	GetSections(limit int, offset int) ([]NewsfeedSectionRepo, error)
	GetSectionsById(id int) (NewsfeedSectionRepo, error)
}
type newsfeedSection struct{}

func NewNewsfeedSectionRepo() *newsfeedSection {
	return &newsfeedSection{}
}
func (n newsfeedSection) GetSections(limit int, offset int) ([]NewsfeedSectionRepo, error) {
	models, err := orm.NewsfeedSection.GetSections(limit, offset)
	if err != nil && err == gorm.ErrRecordNotFound {
		return []NewsfeedSectionRepo{}, nil
	}
	if err != nil {
		return []NewsfeedSectionRepo{}, uerror.InternalError(err, err.Error())
	}
	if len(models) <= 0 {
		return []NewsfeedSectionRepo{}, nil
	}
	sectionIds := make([]int, len(models))
	for i := range models {
		sectionIds[i] = models[i].Id
	}

	return n.populateNewsfeedSections(models), nil
}
func (n newsfeedSection) GetSectionsById(id int) (NewsfeedSectionRepo, error) {
	newsFeedModel, err := orm.NewsfeedSection.GetSectionsById(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return NewsfeedSectionRepo{}, nil
	}
	if err != nil {
		return NewsfeedSectionRepo{}, uerror.InternalError(err, err.Error())
	}
	if newsFeedModel == nil {
		return NewsfeedSectionRepo{}, nil
	}
	return n.populateNewsfeedSection(newsFeedModel), nil
}

func (newsfeedSection) populateNewsfeedSection(input *entities.NewsfeedSection) NewsfeedSectionRepo {
	if input == nil {
		return NewsfeedSectionRepo{}
	}
	resp := NewsfeedSectionRepo{
		Id:        input.Id,
		Name:      input.Name,
		Image:     input.Image,
		Active:    input.Active,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
	return resp
}

func (n newsfeedSection) populateNewsfeedSections(models []entities.NewsfeedSection) []NewsfeedSectionRepo {
	if len(models) <= 0 {
		return []NewsfeedSectionRepo{}
	}
	resp := make([]NewsfeedSectionRepo, len(models))
	for i := range models {
		resp[i] = n.populateNewsfeedSection(&models[i])
	}
	return resp
}
