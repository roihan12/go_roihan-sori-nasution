package repository

import (
	"echo-item/database"
	"echo-item/model"
	"fmt"
)

type ItemRepositoryImpl struct{}

func (i *ItemRepositoryImpl) GetAll(nama string) []model.Item {
	var item []model.Item

	sql := "SELECT * FROM items"
	s := nama

	if s != "" {
		sql = fmt.Sprintf("%s WHERE nama LIKE '%%%s%%' OR deskripsi LIKE '%%%s%%'", sql, s, s)
		// database.DB.Where("nama LIKE ?", nama).Find(&item)
	}

	database.DB.Preload("Category").Raw(sql).Find(&item)

	// database.DB.Preload("Category").Find(&item)

	return item
}

func (i *ItemRepositoryImpl) GetByID(id string) model.Item {
	var item model.Item

	database.DB.Preload("Category").First(&item, "id = ?", id)

	return item
}

func (i *ItemRepositoryImpl) GetByCategoryID(category_id string) []model.Item {
	var item []model.Item

	database.DB.Preload("Category").Find(&item, "category_id = ?", category_id)

	return item
}

func (i *ItemRepositoryImpl) Create(input model.ItemInput) model.Item {
	var newItem model.Item = model.Item{
		Nama:       input.Nama,
		Deskripsi:  input.Deskripsi,
		Stok:       input.Stok,
		Harga:      input.Harga,
		CategoryID: input.CategoryID,
	}

	var createdItem model.Item = model.Item{}

	result := database.DB.Preload("Category").Create(&newItem)

	result.Last(&createdItem)

	return createdItem
}

func (i *ItemRepositoryImpl) Update(id string, input model.ItemInput) model.Item {
	var item model.Item = i.GetByID(id)

	item.Nama = input.Nama
	item.Deskripsi = input.Deskripsi
	item.Stok = input.Stok
	item.Harga = input.Harga
	item.CategoryID = input.CategoryID

	database.DB.Preload("Category").Save(&item)

	return item
}

func (i *ItemRepositoryImpl) Delete(id string) bool {
	var item model.Item = i.GetByID(id)

	result := database.DB.Delete(&item)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
