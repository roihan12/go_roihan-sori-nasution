package category

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(categoryDomain *Domain) Domain
	Update(id string, categoryDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(categoryDomain *Domain) Domain
	Update(id string, categoryDomain *Domain) Domain
	Delete(id string) bool
}
