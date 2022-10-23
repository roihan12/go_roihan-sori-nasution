package category

import (
	"echo-blog/bussiness/category"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name" faker:"name"`
}

func (rec *Category) ToDomain() category.Domain {
	return category.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func FromDomain(domain *category.Domain) *Category {
	return &Category{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
