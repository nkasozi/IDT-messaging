package transport

import (
	"IDT-messaging/core/transports/http/consts"
	"net/http"
	"strings"
)

func authenticationCheckerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		reqToken := request.Header.Get("Authorization")

		if !strings.EqualFold(reqToken, consts.APIKey) {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
