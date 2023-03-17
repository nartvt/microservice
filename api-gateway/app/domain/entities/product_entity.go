package entities

import (
	"time"
)

type Product struct {
	Id        int
	Name      string
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}

type ProductSection struct {
	SectionId int
	ProductId int
}
