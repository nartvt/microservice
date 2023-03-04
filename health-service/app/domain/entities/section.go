package entities

import "time"

type NewsfeedSection struct {
	Id        int
	Name      string
	Active    bool
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
