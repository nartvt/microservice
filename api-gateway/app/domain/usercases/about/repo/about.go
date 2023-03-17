package repo

import (
	"time"

	"gorm.io/gorm"

	"api-gateway/app/domain/entities"
	"api-gateway/app/domain/usercases/about/orm"
	"api-gateway/app/uerror"
)

type AboutRepo struct {
	Id        int
	Name      string
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}
type IAboutRepo interface {
	GetAboutBySectionId(sectionId int, limit int, offset int) ([]AboutRepo, int, error)
}
type aboutRepo struct{}

func NewAboutRepo() *aboutRepo {
	return &aboutRepo{}
}

func (d aboutRepo) GetAboutBySectionId(sectionId int, limit int, offset int) ([]AboutRepo, int, error) {
	abouts, total, err := orm.About.GetAboutBySectionId(sectionId, limit, offset)
	if err != nil && err == gorm.ErrRecordNotFound {
		return []AboutRepo{}, 0, nil
	}
	if err != nil {
		return []AboutRepo{}, 0, uerror.InternalError(err, err.Error())
	}
	return d.populateAbouts(abouts), total, nil
}

func (d aboutRepo) populateAbouts(abouts []entities.About) []AboutRepo {
	if len(abouts) <= 0 {
		return []AboutRepo{}
	}
	resp := make([]AboutRepo, len(abouts))
	for i := range abouts {
		resp[i] = d.populateAbout(abouts[i])
	}
	return resp
}

func (d aboutRepo) populateAbout(about entities.About) AboutRepo {
	return AboutRepo{
		Id:        about.Id,
		Name:      about.Name,
		Image:     about.Image,
		Active:    about.Active,
		CreatedAt: about.CreatedAt,
		UpdatedAt: about.UpdatedAt,
	}
}
