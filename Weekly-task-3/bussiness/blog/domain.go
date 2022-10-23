package blog

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Title        string
	Description  string
	Author       string
	CategoryName string
	CategoryID   uint
}

type Usecase interface {
	GetAll(title string) []Domain
	GetByCategoryID(id string) []Domain
	GetByID(id string) Domain
	Create(blogDomain *Domain) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll(title string) []Domain
	GetByCategoryID(id string) []Domain
	GetByID(id string) Domain
	Create(blogDomain *Domain) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
}
