package projectors

import (
	"context"

	"github.com/vulpes-ferrilata/user-service/view/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProjector interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
}
