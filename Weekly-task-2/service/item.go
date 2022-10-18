package service

import (
	"echo-item/model"
	"echo-item/repository"
)

type ItemService struct {
	Repository repository.ItemRepository
}

func New() ItemService {
	return ItemService{
		Repository: &repository.ItemRepositoryImpl{},
	}
}

func (i *ItemService) GetAll(nama string) []model.Item {
	return i.Repository.GetAll(nama)
}

func (i *ItemService) GetByID(id string) model.Item {
	return i.Repository.GetByID(id)
}

func (i *ItemService) GetByCategoryID(category_id string) []model.Item {
	return i.Repository.GetByCategoryID(category_id)
}

func (i *ItemService) Create(input model.ItemInput) model.Item {
	return i.Repository.Create(input)
}

func (i *ItemService) Update(id string, input model.ItemInput) model.Item {
	return i.Repository.Update(id, input)
}

func (i *ItemService) Delete(id string) bool {
	return i.Repository.Delete(id)
}
