package transport

import (
	"IDT-messaging/core/endpoints/view_models"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting = errors.New("please supply an id in the 'id' field of this url path")
)

func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeSetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req view_models.SetUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return nil, ErrBadRouting
	}

	return view_models.GetUserRequest{Id: id}, nil
}

func decodeListUsersRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return view_models.ListUsersRequest{}, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {

	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(determineHttpResponseStatusCode(err))
	}

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
