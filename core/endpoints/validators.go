package endpoints

import (
	"IDT-messaging/core/endpoints/view_models"
	"IDT-messaging/core/services/consts"
	"errors"
	"fmt"
	"time"
)

func ValidateRequest(r interface{}) error {

	switch r.(type) {
	case view_models.SetUserRequest:
		return validateSetUserRequest(r.(view_models.SetUserRequest))

	case view_models.GetUserRequest:
		return validateGetUserRequest(r.(view_models.GetUserRequest))

	case view_models.ListUsersRequest:
		return validateListUsersRequest(r.(view_models.ListUsersRequest))

	default:
		// by default the validation should fail if it is of an unknown type
		return errors.New("unhandled type sent for validation")
	}
}

func validateListUsersRequest(request view_models.ListUsersRequest) (err error) {
	return
}

func validateGetUserRequest(request view_models.GetUserRequest) (err error) {
	if len(request.Id) <= consts.MinimumLengthOfIDField {
		err = errors.New("please supply a User Id in the 'id' field")
		return
	}
	return
}

func validateSetUserRequest(request view_models.SetUserRequest) (err error) {
	if len(request.Id) <= consts.MinimumLengthOfIDField {
		err = errors.New("please supply a User Id in the 'id' field")
		return
	}

	if len(request.Name) <= consts.MinimumLengthOfNameField {
		err = fmt.Errorf("please supply a User Name with at least %v characters in the 'name' field", consts.MinimumLengthOfNameField)
		return
	}

	unixTime := time.Unix(request.SignUpTime, 0)

	if unixTime.Year() < consts.MinimumYearValueOfSignupTime {
		err = fmt.Errorf("please supply a time whose year is greater than %v in the 'name' field", consts.MinimumYearValueOfSignupTime)
		return
	}

	return
}
