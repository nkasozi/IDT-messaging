package transport

import (
	"IDT-messaging/core/endpoints"
	appConsts "IDT-messaging/core/services/consts"
	"IDT-messaging/core/transports/http/consts"
	"context"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints endpoints.Endpoints) http.Handler {

	r := mux.NewRouter()

	r.Use(jsonContentTypeMiddleware)
	r.Use(authenticationCheckerMiddleware)

	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(encodeError),
	}

	r.Methods(http.MethodPut).Path(consts.IDTSetUserApiRoute).Handler(httpTransport.NewServer(
		endpoints.SetUser,
		decodeSetUserRequest,
		encodeJsonResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path(consts.IDTGetUserApiRoute).Handler(httpTransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeJsonResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path(consts.IDTListUsersApiRoute).Handler(httpTransport.NewServer(
		endpoints.ListUsers,
		decodeListUsersRequest,
		encodeJsonResponse,
		options...,
	))

	return r
}

func determineHttpResponseStatusCode(err error) int {
	switch err {
	case appConsts.InternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
