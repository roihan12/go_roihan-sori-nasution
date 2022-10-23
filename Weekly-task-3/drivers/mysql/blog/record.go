package blog

import (
	blogUsecase "echo-blog/bussiness/blog"
	"echo-blog/drivers/mysql/category"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID          uint              `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `json:"deleted_at"`
	Title       string            `json:"title" faker:"word"`
	Description string            `json:"description " faker:"sentence"`
	Author      string            `json:"author" faker:"word"`
	Category    category.Category `json:"category"`
	CategoryID  uint              `json:"category_id"`
}

func FromDomain(domain *blogUsecase.Domain) *Blog {
	return &Blog{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		Author:      domain.Author,
		CategoryID:  domain.CategoryID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func (rec *Blog) ToDomain() blogUsecase.Domain {
	return blogUsecase.Domain{
		ID:           rec.ID,
		Title:        rec.Title,
		Description:  rec.Description,
		Author:       rec.Author,
		CategoryName: rec.Category.Name,
		CategoryID:   rec.Category.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}
