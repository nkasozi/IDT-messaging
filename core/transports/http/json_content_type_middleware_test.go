package transport

import (
	"IDT-messaging/core/transports/http/consts"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonContentTypeMiddleware(t *testing.T) {
	t.Run("jsonContentTypeMiddleware should add Json Content Type header", func(t *testing.T) {

		// create a handler to use as "next" which will verify the request
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := w.Header().Get(consts.ContentTypeHeaderName)

			if val == "" {
				t.Error("Missing Content Type Header")
				return
			}
		})

		// create the handler to test, using our custom "next" handler
		handlerToTest := jsonContentTypeMiddleware(nextHandler)

		// create a mock request to use
		req := httptest.NewRequest("GET", "http://testing", nil)

		// call the handler using a mock response recorder (we'll not use that anyway)
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	})
}
