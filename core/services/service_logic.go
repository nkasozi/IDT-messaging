package services

import (
	"IDT-messaging/core/services/models"
	"context"
	"github.com/go-kit/kit/log"
)

type service struct {
	logger    log.Logger
	usersRepo UsersRepo
}

func (s service) SetUser(ctx context.Context, request models.User) (response models.User, err error) {

	err = s.usersRepo.SaveUser(request)

	if err != nil {
		return
	}

	response = request
	return
}

func (s service) GetUser(ctx context.Context, id string) (response models.User, err error) {

	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return
	}

	response = user
	return
}

func (s service) ListUsers(ctx context.Context) (users []models.User, err error) {
	users, err = s.usersRepo.GetAllUsers()
	return
}

func NewService(logger log.Logger, usersRepo UsersRepo) Service {
	return &service{
		logger:    logger,
		usersRepo: usersRepo,
	}
}
