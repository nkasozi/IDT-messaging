package transport

import (
	"IDT-messaging/core/endpoints/view_models"
	"IDT-messaging/core/transports/http/consts"
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeSetUserRequest(t *testing.T) {
	t.Run("given valid request should return no error", func(t *testing.T) {
		ctx := context.Background()
		url := "http://test-url.com/v1/users"
		expectedDecodedRequest := view_models.SetUserRequest{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}
		body := []byte(`{
				"id":"test-user-1",
				"name":"test-user-1",
				"signUpTime":6000
	     }`)
		httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		httpRequest.Header.Set(consts.ContentTypeHeaderName, consts.JsonContentTypeValue)
		decodedRequest, err := decodeSetUserRequest(ctx, httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, expectedDecodedRequest, decodedRequest)
	})

	t.Run("given an invalid request should return an error", func(t *testing.T) {
		ctx := context.Background()
		url := "http://test-url.com/v1/users"
		body := []byte(`{
				"test-invalid-json": 123,
	     }`)
		httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		httpRequest.Header.Set(consts.ContentTypeHeaderName, consts.JsonContentTypeValue)
		_, err = decodeSetUserRequest(ctx, httpRequest)
		assert.Error(t, err)
	})
}

func TestEncodeJsonResponse(t *testing.T) {
	t.Run("given valid response should return no error", func(t *testing.T) {
		ctx := context.Background()
		writer := new(httptest.ResponseRecorder)
		response := view_models.SetUserResponse{
			Id:         "test-user-1",
			Name:       "test-user-1",
			SignUpTime: 6000,
		}
		err := encodeJsonResponse(ctx, writer, response)
		assert.NoError(t, err)
	})
}
