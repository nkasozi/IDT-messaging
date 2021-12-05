package endpoints

import (
	"IDT-messaging/core/endpoints/view_models"
	"IDT-messaging/core/services/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSetUserRequest(t *testing.T) {
	t.Run("should return no error given valid SetUserRequest", func(t *testing.T) {
		request := view_models.SetUserRequest{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		expected := models.User{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		actual, err := ToSetUserServiceRequest(request)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return an error given invalid request", func(t *testing.T) {
		request := models.User{}
		_, err := ToSetUserServiceRequest(request)
		assert.Error(t, err)
	})
}

func TestToGetUserRequest(t *testing.T) {
	t.Run("should return no error given valid GetUserRequest", func(t *testing.T) {
		request := view_models.GetUserRequest{
			Id: "test-user-1",
		}

		expected := "test-user-1"

		actual, err := ToGetUserServiceRequest(request)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return an error given invalid request", func(t *testing.T) {
		request := models.User{}
		_, err := ToGetUserServiceRequest(request)
		assert.Error(t, err)
	})
}
