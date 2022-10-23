package category

import (
	"echo-blog/bussiness/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) category.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll() []category.Domain {
	var rec []Category

	cr.conn.Find(&rec)

	categoryDomain := []category.Domain{}

	for _, category := range rec {
		categoryDomain = append(categoryDomain, category.ToDomain())
	}

	return categoryDomain
}

func (cr *categoryRepository) GetByID(id string) category.Domain {
	var category Category

	cr.conn.First(&category, "id = ?", id)

	return category.ToDomain()
}

func (cr *categoryRepository) Create(categoryDomain *category.Domain) category.Domain {
	rec := FromDomain(categoryDomain)

	result := cr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (cr *categoryRepository) Update(id string, categoryDomain *category.Domain) category.Domain {
	var category category.Domain = cr.GetByID(id)

	updatedCategory := FromDomain(&category)

	updatedCategory.Name = categoryDomain.Name

	cr.conn.Save(&updatedCategory)

	return updatedCategory.ToDomain()
}

func (cr *categoryRepository) Delete(id string) bool {
	var category category.Domain = cr.GetByID(id)

	deletedCategory := FromDomain(&category)

	result := cr.conn.Unscoped().Delete(&deletedCategory)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
