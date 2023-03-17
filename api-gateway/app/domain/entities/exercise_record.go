package entities

import "time"

type UserExercise struct {
	Id             int
	UserId         int
	AtTime         int
	Description    string
	CaloriesBurned int
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
