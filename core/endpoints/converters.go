package endpoints

import (
	view_models "IDT-messaging/core/endpoints/view_models"
	"IDT-messaging/core/services/models"
	"errors"
)

func ToSetUserServiceRequest(request interface{}) (req models.User, err error) {
	setUserReq, ok := request.(view_models.SetUserRequest)
	if !ok {
		err = errors.New("unsupported request supplied")
		return
	}
	req = models.User{
		Id:         setUserReq.Id,
		Name:       setUserReq.Name,
		SignUpTime: setUserReq.SignUpTime,
	}
	return
}

func ToGetUserServiceRequest(request interface{}) (req string, err error) {
	getUserReq, ok := request.(view_models.GetUserRequest)
	if !ok {
		err = errors.New("unsupported request supplied")
		return
	}
	req = getUserReq.Id
	return
}

func ToListUsersResponse(users []models.User) (resp view_models.ListUsersResponse) {
	for _, user := range users {
		viewModelUser := view_models.User{
			Id:         user.Id,
			Name:       user.Name,
			SignUpTime: user.SignUpTime,
		}

		resp.Users = append(resp.Users, viewModelUser)
	}

	return
}
