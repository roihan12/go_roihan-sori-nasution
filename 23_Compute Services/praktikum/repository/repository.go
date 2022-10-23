package repository

import "echo-item/model"

type ItemRepository interface {
	GetAll(nama string) []model.Item
	GetByID(id string) model.Item
	GetByCategoryID(category_id string) []model.Item
	Create(input model.ItemInput) model.Item
	Update(id string, input model.ItemInput) model.Item
	Delete(id string) bool
}

type CategoryRepository interface {
	GetAll() []model.Category
	GetByID(id string) model.Category
	Create(input model.CategoryInput) model.Category
	Update(id string, input model.CategoryInput) model.Category
	Delete(id string) bool
}
