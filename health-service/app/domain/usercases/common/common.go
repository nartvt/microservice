package common

import (
	"gorm.io/gorm"

	"health-service/app/infra/db"
)

const (
	NewsfeedSectionTypeTopPage        = "top_page"
	NewsfeedSectionTypePersonalRecord = "personal_record"
	NewsfeedSectionTypeAbout          = "about_recommendations"
)

func BeginTx() *gorm.DB {
	return db.Postgres.Begin()
}

func RecoveryTx(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
		panic(err)
	}
}
