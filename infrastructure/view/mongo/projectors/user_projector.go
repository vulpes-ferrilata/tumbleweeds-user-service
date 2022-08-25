package projectors

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/user-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/user-service/infrastructure/view/mongo/mappers"
	"github.com/vulpes-ferrilata/user-service/view/models"
	"github.com/vulpes-ferrilata/user-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserProjector(db *mongo.Database) projectors.UserProjector {
	return &userProjector{
		userCollection: db.Collection("users"),
	}
}

type userProjector struct {
	userCollection *mongo.Collection
}

func (u userProjector) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	userDocument := &documents.User{}

	filter := bson.M{"_id": id}

	if err := u.userCollection.FindOne(ctx, filter).Decode(userDocument); err != nil {
		return nil, errors.WithStack(err)
	}

	user := mappers.ToUserView(userDocument)

	return user, nil
}
