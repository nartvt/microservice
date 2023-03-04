package entities

import "time"

type User struct {
	Id          int
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	FirstName   string
	LastName    string
	FullName    string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
