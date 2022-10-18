package service

import (
	"echo-item/model"
	"echo-item/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func NewCategoryService() CategoryService {
	return CategoryService{
		Repository: &repository.CategoryRepositoryImpl{},
	}
}

func (c *CategoryService) GetAll() []model.Category {
	return c.Repository.GetAll()
}

func (c *CategoryService) GetByID(id string) model.Category {
	return c.Repository.GetByID(id)
}

func (c *CategoryService) Create(input model.CategoryInput) model.Category {
	return c.Repository.Create(input)
}

func (c *CategoryService) Update(id string, input model.CategoryInput) model.Category {
	return c.Repository.Update(id, input)
}

func (c *CategoryService) Delete(id string) bool {
	return c.Repository.Delete(id)
}
