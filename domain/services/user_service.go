package services

import (
	"context"

	"github.com/VulpesFerrilata/user-service/domain/models"
	"github.com/google/uuid"
)

type UserService interface {
	NewUser(ctx context.Context, id uuid.UUID, displayName string) (*models.User, error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct{}

func (u userService) NewUser(ctx context.Context, id uuid.UUID, displayName string) (*models.User, error) {
	user := models.NewUser(id, displayName, 0)

	return user, nil
}
