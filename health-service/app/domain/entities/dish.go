package entities

import (
	"time"
)

type Dish struct {
	Id        int
	Name      string
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}

type SectionDish struct {
	SectionId int
	DishId    int
}
