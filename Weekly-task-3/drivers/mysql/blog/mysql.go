package blog

import (
	"echo-blog/bussiness/blog"
	"fmt"

	"gorm.io/gorm"
)

type blogRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) blog.Repository {
	return &blogRepository{
		conn: conn,
	}
}

func (br *blogRepository) GetByCategoryID(category_id string) []blog.Domain {
	var rec []Blog

	br.conn.Preload("Category").Find(&rec, "category_id = ?", category_id)

	blogDomain := []blog.Domain{}

	for _, blog := range rec {
		blogDomain = append(blogDomain, blog.ToDomain())
	}

	return blogDomain
}

func (br *blogRepository) GetAll(title string) []blog.Domain {
	var rec []Blog

	sql := "SELECT * FROM blogs"
	s := title

	if s != "" {
		sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
	}

	// br.conn.Preload("Category").Find(&rec)
	br.conn.Preload("Category").Raw(sql).Find(&rec)

	blogDomain := []blog.Domain{}

	for _, blog := range rec {
		blogDomain = append(blogDomain, blog.ToDomain())
	}

	return blogDomain
}

func (br *blogRepository) GetByID(id string) blog.Domain {
	var blog Blog

	br.conn.Preload("Category").First(&blog, "id = ?", id)

	return blog.ToDomain()
}

func (br *blogRepository) Create(blogDomain *blog.Domain) blog.Domain {
	rec := FromDomain(blogDomain)

	result := br.conn.Preload("Category").Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (br *blogRepository) Update(id string, blogDomain *blog.Domain) blog.Domain {
	var blog blog.Domain = br.GetByID(id)

	updatedBlog := FromDomain(&blog)

	updatedBlog.Title = blogDomain.Title
	updatedBlog.Description = blogDomain.Description
	updatedBlog.Author = blogDomain.Author
	updatedBlog.CategoryID = blogDomain.CategoryID

	br.conn.Save(&updatedBlog)

	return updatedBlog.ToDomain()
}

func (br *blogRepository) Delete(id string) bool {
	var blog blog.Domain = br.GetByID(id)

	deletedBlog := FromDomain(&blog)

	result := br.conn.Delete(&deletedBlog)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
