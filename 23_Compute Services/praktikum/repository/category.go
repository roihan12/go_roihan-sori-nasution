package repository

import (
	"echo-item/database"
	"echo-item/model"
)

type CategoryRepositoryImpl struct{}

func (c *CategoryRepositoryImpl) GetAll() []model.Category {
	var categories []model.Category

	database.DB.Preload("Category").Find(&categories)

	return categories
}

func (c *CategoryRepositoryImpl) GetByID(id string) model.Category {
	var category model.Category

	database.DB.First(&category, "id = ?", id)

	return category
}

func (c *CategoryRepositoryImpl) Create(input model.CategoryInput) model.Category {
	var newCategory model.Category = model.Category{
		Nama: input.Nama,
	}

	var createdCategory model.Category = model.Category{}

	result := database.DB.Create(&newCategory)

	result.Last(&createdCategory)

	return createdCategory
}

func (c *CategoryRepositoryImpl) Update(id string, input model.CategoryInput) model.Category {
	var category model.Category = c.GetByID(id)

	category.Nama = input.Nama

	database.DB.Save(&category)

	return category
}

func (c *CategoryRepositoryImpl) Delete(id string) bool {
	var category model.Category = c.GetByID(id)

	result := database.DB.Unscoped().Delete(&category)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
