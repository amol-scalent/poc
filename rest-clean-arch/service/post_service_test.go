package service

import (
	"testing"

	"github.com/amolasg/go-projects/clean-arch-example/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := m.Called()

	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (m *MockRepository) FindAll() ([]entity.Post, error) {
	args := m.Called()

	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

// TODO:
func (m *MockRepository) Delete(post *entity.Post) error {
	return nil
}
func TestFindAll(t *testing.T) {

	var id int64 = 1
	mockRepo := new(MockRepository)

	post := entity.Post{ID: id, Title: "ABC", Text: "Text"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, id, result[0].ID)
	assert.Equal(t, "ABC", result[0].Title)
	assert.Equal(t, "Text", result[0].Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "test"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	var id int64 = 1

	post := entity.Post{ID: id, Title: "ABC", Text: "Text"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(post)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, id, result.ID)
	assert.Equal(t, "ABC", result.Title)
	assert.Equal(t, "Text", result.Text)
	assert.Nil(t, err)
}
