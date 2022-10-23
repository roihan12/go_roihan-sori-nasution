package blog_test

import (
	"echo-blog/bussiness/blog"
	_blogMock "echo-blog/bussiness/blog/mocks"
	"echo-blog/bussiness/category"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	blogRepository _blogMock.Repository
	blogService    blog.Usecase

	blogDomain blog.Domain
)

func TestMain(m *testing.M) {
	blogService = blog.NewBlogUsecase(&blogRepository)

	categoryDomain := category.Domain{
		Name: "test category",
	}

	blogDomain = blog.Domain{
		Title:       "title",
		Description: "my descriptions",
		Author:      "my author",
		CategoryID:  categoryDomain.ID,
	}

	m.Run()
}

// func TestGetAll(t *testing.T) {
// 	t.Run("Get All | Valid", func(t *testing.T) {
// 		blogRepository.On("GetAll").Return([]blog.Domain{blogDomain}).Once()

// 		result := blogService.GetAll()

// 		assert.Equal(t, 1, len(result))
// 	})

// 	t.Run("Get All | InValid", func(t *testing.T) {
// 		blogRepository.On("GetAll").Return([]blog.Domain{}).Once()

// 		result := blogService.GetAll()

// 		assert.Equal(t, 0, len(result))
// 	})
// }

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		blogRepository.On("GetByID", "1").Return(blogDomain).Once()

		result := blogService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		blogRepository.On("GetByID", "-1").Return(blog.Domain{}).Once()

		result := blogService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestGetByCategoryID(t *testing.T) {
	t.Run("Get By CategoryID | Valid", func(t *testing.T) {
		blogRepository.On("GetByCategoryID", "1").Return([]blog.Domain{blogDomain}).Once()

		result := blogService.GetByCategoryID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By CategoryID | InValid", func(t *testing.T) {
		blogRepository.On("GetByCategoryID", "-1").Return([]blog.Domain{blogDomain}).Once()

		result := blogService.GetByCategoryID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		blogRepository.On("Create", &blogDomain).Return(blogDomain).Once()

		result := blogService.Create(&blogDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		blogRepository.On("Create", &blog.Domain{}).Return(blog.Domain{}).Once()

		result := blogService.Create(&blog.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		blogRepository.On("Update", "1", &blogDomain).Return(blogDomain).Once()

		result := blogService.Update("1", &blogDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		blogRepository.On("Update", "1", &blog.Domain{}).Return(blog.Domain{}).Once()

		result := blogService.Update("1", &blog.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		blogRepository.On("Delete", "1").Return(true).Once()

		result := blogService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		blogRepository.On("Delete", "-1").Return(false).Once()

		result := blogService.Delete("-1")

		assert.False(t, result)
	})
}
