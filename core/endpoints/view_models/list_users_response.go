package view_models

type ListUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	SignUpTime int64  `json:"signupTime"`
}
