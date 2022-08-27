package service

import (
	"fmt"
	"testing"

	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	fmt.Println("args", args)
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *mockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	fmt.Println("args", args)
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *mockRepository) FindByID(id string) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	var expectedID int64 = 1
	post := entity.Post{ID: expectedID, Title: "title", Text: "text"}
	// setup the expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	// mock assertion: behavioral
	mockRepo.AssertExpectations(t)

	// data assertion
	assert.Equal(t, "title", result[0].Title)
	assert.Equal(t, "text", result[0].Text)
	assert.Equal(t, expectedID, result[0].ID)
}

func TestSave(t *testing.T) {
	mockRepo := new(mockRepository)
	post := entity.Post{Title: "title", Text: "text"}
	// setup the expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.Create(&post)
	// mock assertion: behavioral
	mockRepo.AssertExpectations(t)

	// data assertion
	assert.Equal(t, "title", result.Title)
	assert.Equal(t, "text", result.Text)
	assert.NotNil(t, post.ID)
}

func TestValidateEmptyPost(t *testing.T) {
	// passing nil to NewPostService because
	// its not using any repo for operation...
	testService := NewPostService(nil)
	err := testService.Validate(nil)

	assert.NotNil(t, err)
	// expected, actual
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {

	post := entity.Post{ID: 1, Title: "", Text: "Text"}

	testService := NewPostService(nil)
	err := testService.Validate(&post)

	assert.NotNil(t, err)
	// expected, actual
	assert.Equal(t, "the post's title is empty", err.Error())

}
