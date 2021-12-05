package repository

import (
	"IDT-messaging/core/services"
	"IDT-messaging/core/services/consts"
	"IDT-messaging/core/services/models"
)

type usersRepo struct {
	Users map[string]*models.User
}

func (u usersRepo) SaveUser(user models.User) (err error) {
	if u.Users == nil {
		return consts.InternalServerError
	}
	u.Users[user.Id] = &user
	return
}

func (u usersRepo) GetUserById(id string) (user models.User, err error) {

	if u.Users == nil {
		err = consts.InternalServerError
		return
	}

	userPtr, found := u.Users[id]

	if found {
		user = *userPtr
		return
	}

	err = consts.UserNotFoundError
	return
}

func (u usersRepo) GetAllUsers() (users []models.User, err error) {
	if u.Users == nil {
		err = consts.InternalServerError
		return
	}

	for _, user := range u.Users {
		users = append(users, *user)
	}
	return
}

func NewUsersRepo() services.UsersRepo {
	return &usersRepo{
		Users: make(map[string]*models.User, 0),
	}
}
