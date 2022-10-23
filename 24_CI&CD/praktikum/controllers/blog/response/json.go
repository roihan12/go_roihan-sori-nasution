package response

import (
	"echo-blog/bussiness/blog"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Author       string         `json:"author"`
	CategoryName string         `json:"category_name"`
	CategoryID   uint           `json:"category_id"`
}

func FromDomain(domain blog.Domain) Blog {
	return Blog{
		ID:           domain.ID,
		Title:        domain.Title,
		Description:  domain.Description,
		Author:       domain.Author,
		CategoryName: domain.CategoryName,
		CategoryID:   domain.CategoryID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}
