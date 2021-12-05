package endpoints

import (
	"IDT-messaging/core/services"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SetUser   endpoint.Endpoint
	GetUser   endpoint.Endpoint
	ListUsers endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		SetUser:   makeSetUserEndpoint(s),
		GetUser:   makeGetUserEndpoint(s),
		ListUsers: makeListUsersEndpoint(s),
	}
}

func makeSetUserEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		err = ValidateRequest(request)

		if err != nil {
			return
		}

		req, err := ToSetUserServiceRequest(request)

		if err != nil {
			return
		}

		response, err = s.SetUser(ctx, req)
		return
	}
}

func makeGetUserEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		err = ValidateRequest(request)

		if err != nil {
			return
		}

		req, err := ToGetUserServiceRequest(request)

		if err != nil {
			return
		}

		response, err = s.GetUser(ctx, req)
		return
	}
}

func makeListUsersEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		err = ValidateRequest(request)

		if err != nil {
			return
		}

		allUsers, err := s.ListUsers(ctx)

		response = ToListUsersResponse(allUsers)

		return
	}
}
