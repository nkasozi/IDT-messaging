package services

import (
	"IDT-messaging/core/endpoints/view_models"
	"IDT-messaging/core/services/mocks"
	"IDT-messaging/core/services/models"
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetUser(t *testing.T) {

	t.Run("SetUser should return no error if the request is valid", func(t *testing.T) {
		service, mockUsersRepo := setUp(t)

		request := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		mockUsersRepo.EXPECT().SaveUser(request).Return(nil)

		expected := view_models.SetUserResponse{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		response, err := service.SetUser(context.Background(), request)
		assert.NoError(t, err)
		assert.EqualValues(t, expected, response)
	})

	t.Run("SetUser should return an error if the usersRepo returns an error", func(t *testing.T) {
		service, mockUsersRepo := setUp(t)

		request := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		mockUsersRepo.EXPECT().SaveUser(request).Return(errors.New("dummy error from users repo"))

		_, err := service.SetUser(context.Background(), request)
		assert.Error(t, err)
	})
}

func TestGetUser(t *testing.T) {

	t.Run("GetUser should return no error if the request is valid", func(t *testing.T) {
		service, mockUsersRepo := setUp(t)

		request := "test-user-1"
		dummyUser := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		mockUsersRepo.EXPECT().GetUserById(request).Return(dummyUser, nil)

		expected := view_models.GetUserResponse{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		response, err := service.GetUser(context.Background(), request)
		assert.NoError(t, err)
		assert.EqualValues(t, expected, response)
	})

	t.Run("GetUser should return an error if the usersRepo returns an error", func(t *testing.T) {
		service, mockUsersRepo := setUp(t)

		request := "test-user-1"

		mockUsersRepo.EXPECT().GetUserById(request).Return(models.User{}, errors.New("dummy error from users repo"))

		_, err := service.GetUser(context.Background(), request)
		assert.Error(t, err)
	})
}

func TestListUsers(t *testing.T) {

	t.Run("ListUsers should return no error if the request is valid", func(t *testing.T) {
		service, mockUsersRepo := setUp(t)

		dummyUsers := []models.User{
			{
				Id:         "test-user-1",
				Name:       "test-user-1",
				SignUpTime: 6000,
			},
		}

		mockUsersRepo.EXPECT().GetAllUsers().Return(dummyUsers, nil)

		expected := []models.User{
			{
				Id:         "test-user-1",
				Name:       "test-user-1",
				SignUpTime: 6000,
			},
		}

		response, err := service.ListUsers(context.Background())
		assert.NoError(t, err)
		assert.EqualValues(t, expected, response)
	})

	t.Run("ListUsers should return an error if the usersRepo returns an error", func(t *testing.T) {
		service, mockUsersRepo := setUp(t)

		mockUsersRepo.EXPECT().GetAllUsers().Return([]models.User{}, errors.New("dummy error from users repo"))

		_, err := service.ListUsers(context.Background())
		assert.Error(t, err)
	})
}

func setUp(t *testing.T) (Service, *mocks.MockUsersRepo) {
	ctrl := gomock.NewController(t)
	mockUsersRepo := mocks.NewMockUsersRepo(ctrl)
	return NewService(log.NewNopLogger(), mockUsersRepo), mockUsersRepo
}
