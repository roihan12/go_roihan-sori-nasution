package drivers

import (
	categoryDomain "echo-blog/bussiness/category"
	categoryDB "echo-blog/drivers/mysql/category"

	blogDomain "echo-blog/bussiness/blog"
	blogDB "echo-blog/drivers/mysql/blog"

	userDomain "echo-blog/bussiness/users"
	userDB "echo-blog/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewBlogRepository(conn *gorm.DB) blogDomain.Repository {
	return blogDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
