package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Nama       string    `json:"nama"`
	Deskripsi  string    `json:"deskripsi"`
	Stok       int       `json:"stok"`
	Harga      float64   `json:"harga"`
	Category   Category  `json:"category"`
	CategoryID uint      `json:"category_id"`
}

func (Item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	Item.ID = uuid.New()
	return
}
