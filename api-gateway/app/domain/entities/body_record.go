package entities

import "time"

type UserBodyRecord struct {
	Id         int
	UserId     int
	Weight     float32
	Height     int
	Percentage float32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
