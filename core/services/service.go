package services

import (
	"IDT-messaging/core/services/models"
	"context"
)

//go:generate mockgen -destination=mocks/mock_service.go -package=mocks . Service
type Service interface {
	SetUser(ctx context.Context, request models.User) (response models.User, err error)
	GetUser(ctx context.Context, id string) (response models.User, err error)
	ListUsers(ctx context.Context) (response []models.User, err error)
}

//go:generate mockgen -destination=mocks/mock_usersRepo.go -package=mocks . UsersRepo
type UsersRepo interface {
	SaveUser(user models.User) (err error)
	GetUserById(id string) (user models.User, err error)
	GetAllUsers() (users []models.User, err error)
}
