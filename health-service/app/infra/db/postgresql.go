package db

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	"health-service/app/config"
)

var Postgres *gorm.DB

func InitPostgres() {
	conf := config.Config
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Posgres.UserName,
		conf.Posgres.Password,
		conf.Posgres.Host,
		conf.Posgres.Port,
		conf.Posgres.Database,
	)
	var err error
	Postgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func ClosePostgres() {
	if db, _ := Postgres.DB(); db != nil {
		if err := db.Close(); err != nil {
			fmt.Println("[ERROR] Cannot close mysql connection, err:", err)
		}
	}
}
