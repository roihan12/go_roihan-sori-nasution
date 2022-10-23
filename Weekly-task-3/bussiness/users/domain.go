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
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) string
}

type Repository interface {
	Register(userDomain *Domain) Domain
	GetByEmail(userDomain *Domain) Domain
}
