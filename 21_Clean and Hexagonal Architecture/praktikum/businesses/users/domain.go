package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
}

type Usecase interface {
	CreateUser(userDomain *Domain) Domain
	Login(userDomain *Domain) string
	GetAllusers() []Domain
}

type Repository interface {
	CreateUser(userDomain *Domain) Domain
	GetByEmail(userDomain *Domain) Domain
	GetAllusers() []Domain
}
