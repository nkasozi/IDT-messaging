package repository

import (
	"IDT-messaging/core/services"
	"IDT-messaging/core/services/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveUser(t *testing.T) {
	t.Run("should return no error given valid SaveUser Request", func(t *testing.T) {

		usersRepo := setUp()

		request := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		err := usersRepo.SaveUser(request)
		assert.NoError(t, err)
	})

	t.Run("should return an error given any other error on save", func(t *testing.T) {

		usersRepo := usersRepo{}

		request := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		err := usersRepo.SaveUser(request)
		assert.Error(t, err)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("should return no error given valid SaveUser Request", func(t *testing.T) {

		usersRepo := usersRepo{
			Users: map[string]*models.User{
				"test-user-1": {
					Id:         "test-user-1",
					Name:       "test-user-1",
					SignUpTime: 6000,
				},
			},
		}

		request := "test-user-1"

		expected := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		user, err := usersRepo.GetUserById(request)
		assert.NoError(t, err)
		assert.EqualValues(t, expected, user)
	})

	t.Run("should return an error given any other error on get", func(t *testing.T) {
		usersRepo := usersRepo{}
		request := "test-user-1"
		_, err := usersRepo.GetUserById(request)
		assert.Error(t, err)
	})
}

func TestListUsers(t *testing.T) {
	t.Run("should return no error given valid ListUsers Request", func(t *testing.T) {

		usersRepo := usersRepo{
			Users: map[string]*models.User{
				"test-user-1": {
					Id:         "test-user-1",
					Name:       "test-user-1",
					SignUpTime: 6000,
				},
			},
		}

		expected := []models.User{
			{
				Id:         "test-user-1",
				Name:       "test-user-1",
				SignUpTime: 6000,
			},
		}

		users, err := usersRepo.GetAllUsers()
		assert.NoError(t, err)
		assert.EqualValues(t, expected, users)
	})

	t.Run("should return an error given any other error on get", func(t *testing.T) {
		usersRepo := usersRepo{}
		_, err := usersRepo.GetAllUsers()
		assert.Error(t, err)
	})
}

func setUp() services.UsersRepo {
	return NewUsersRepo()
}
