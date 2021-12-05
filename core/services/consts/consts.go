package consts

import "errors"

const (
	MinimumLengthOfIDField       = 0
	MinimumLengthOfNameField     = 2
	MinimumYearValueOfSignupTime = 1850
)

var UserNotFoundError = errors.New("user not found")
var InternalServerError = errors.New("internal server error")
