package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/user-service/domain/models"
)

type UserRepository interface {
	Insert(ctx context.Context, user *models.User) error
}
