package usecase_test

import (
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct{}

func (mockRepo *mockUserRepository) GetByName(name string) (*model.User, error) {
	var mockUser model.User
	faker.FakeData(&mockUser)

	return &mockUser, nil
}

func (mockRepo *mockUserRepository) Create(User *model.User) error {
	return nil
}

func TestUserGetByName(t *testing.T) {
	mockRepo := &mockUserRepository{}
	userUsecase := usecase.NewUserUsecase(mockRepo)

	User, err := userUsecase.GetByName("test")
	assert.NoError(t, err)
	assert.NotNil(t, User)
}

func TestUserCreate(t *testing.T) {
	mockRepo := &mockUserRepository{}
	userUsecase := usecase.NewUserUsecase(mockRepo)

	User := model.User{}
	faker.FakeData(&User)

	err := userUsecase.Create(&User)
	assert.NoError(t, err)
}
