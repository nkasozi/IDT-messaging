package view_models

type SetUserRequest struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	SignUpTime int64  `json:"signupTime"`
}
