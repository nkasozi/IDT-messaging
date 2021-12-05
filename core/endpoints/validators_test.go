package endpoints

import (
	"IDT-messaging/core/endpoints/view_models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateSetUserRequest(t *testing.T) {
	t.Run("should return no error given valid SetUserRequest", func(t *testing.T) {
		request := view_models.SetUserRequest{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		err := ValidateRequest(request)
		assert.NoError(t, err)
	})

	t.Run("should return an error given request with empty ID", func(t *testing.T) {
		request := view_models.SetUserRequest{
			Name:       "test-user-1",
			SignUpTime: 6000,
		}

		err := ValidateRequest(request)
		assert.Error(t, err)
	})

	t.Run("should return an error given request with name less than 2 characters", func(t *testing.T) {
		request := view_models.SetUserRequest{
			Id:         "test-user-1",
			Name:       "t",
			SignUpTime: 6000,
		}

		err := ValidateRequest(request)
		assert.Error(t, err)
	})

	t.Run("should return an error given request with SignupTime less than 1850", func(t *testing.T) {
		request := view_models.SetUserRequest{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: -3818361600,
		}

		err := ValidateRequest(request)
		assert.Error(t, err)
	})
}

func TestValidateGetUserRequest(t *testing.T) {
	t.Run("should return no error given valid GetUserRequest", func(t *testing.T) {
		request := view_models.GetUserRequest{
			Id: "test-user-1",
		}

		err := ValidateRequest(request)
		assert.NoError(t, err)
	})

	t.Run("should return an error given request with empty ID", func(t *testing.T) {
		request := view_models.GetUserRequest{}

		err := ValidateRequest(request)
		assert.Error(t, err)
	})
}
