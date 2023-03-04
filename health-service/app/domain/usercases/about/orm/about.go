package orm

import (
	"health-service/app/domain/entities"
	"health-service/app/domain/usercases/common"
	"health-service/app/infra/db"
)

type IAbout interface {
	GetAboutBySectionId(sectionId int, limit int, offset int) ([]entities.About, int, error)
}
type about struct{}

var About IAbout

func init() {
	About = about{}
}
func (a about) GetAboutBySectionId(sectionId int, limit int, offset int) ([]entities.About, int, error) {
	resp := []entities.About{}
	total := int64(0)
	err := db.Postgres.Model(&entities.About{}).
		Joins("JOIN sections_about ON abouts.id = sections_about.about_id").
		Joins("JOIN newsfeed_sections ON sections_about.section_id = newsfeed_sections.id").
		Where("newsfeed_sections.active = TRUE").
		Where("abouts.active = TRUE").
		Where("type = ?", common.NewsfeedSectionTypeAbout).
		Where("newsfeed_sections.id = ?", sectionId).
		Limit(limit).
		Offset(offset).
		Order("id DESC").
		Count(&total).
		Find(&resp).Error
	return resp, int(total), err
}
