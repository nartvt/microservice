package entities

import "time"

type UserDiary struct {
	Id          int
	UserId      int
	AtTime      int
	Description string
	Calories    int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
