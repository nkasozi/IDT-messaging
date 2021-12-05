package transport

import (
	"IDT-messaging/core/transports/http/consts"
	"net/http"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add(consts.ContentTypeHeaderName, consts.JsonContentTypeValue)
		next.ServeHTTP(writer, request)
	})
}
