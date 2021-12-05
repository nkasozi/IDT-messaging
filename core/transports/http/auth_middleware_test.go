package transport

import (
	"IDT-messaging/core/transports/http/consts"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_authenticationCheckerMiddleware(t *testing.T) {
	t.Run("authenticationCheckerMiddleware should fail requests without Authorization Header", func(t *testing.T) {

		// create a handler to use as "next" which will verify the request
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("auth middleware should not call next handler if Auth Header is missing")
			return
		})

		// create the handler to test, using our custom "next" handler
		handlerToTest := authenticationCheckerMiddleware(nextHandler)

		// create a mock request to use
		req := httptest.NewRequest("GET", "http://testing", nil)

		// call the handler using a mock response recorder (we'll not use that anyway)
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	})

	t.Run("authenticationCheckerMiddleware should fail requests with wrong Authorization Header Value", func(t *testing.T) {

		// create a handler to use as "next" which will verify the request
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("auth middleware should not call next handler if Auth Header is missing")
			return
		})

		// create the handler to test, using our custom "next" handler
		handlerToTest := authenticationCheckerMiddleware(nextHandler)

		// create a mock request to use
		req := httptest.NewRequest("GET", "http://testing", nil)
		req.Header.Add("Authorization", "random-incorrect-key")

		// call the handler using a mock response recorder (we'll not use that anyway)
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	})

	t.Run("authenticationCheckerMiddleware should pass requests with correct Authorization Header", func(t *testing.T) {

		// create a handler to use as "next" which will verify the request
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			return
		})

		// create the handler to test, using our custom "next" handler
		handlerToTest := authenticationCheckerMiddleware(nextHandler)

		// create a mock request to use
		req := httptest.NewRequest("GET", "http://testing", nil)
		req.Header.Add("Authorization", consts.APIKey)

		// call the handler using a mock response recorder (we'll not use that anyway)
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	})
}
