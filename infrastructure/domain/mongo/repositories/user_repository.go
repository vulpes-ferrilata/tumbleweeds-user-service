package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/user-service/domain/models"
	"github.com/vulpes-ferrilata/user-service/domain/repositories"
	"github.com/vulpes-ferrilata/user-service/infrastructure/domain/mongo/mappers"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) repositories.UserRepository {
	return &userRepository{
		userCollection: db.Collection("users"),
	}
}

type userRepository struct {
	userCollection *mongo.Collection
}

func (u userRepository) Insert(ctx context.Context, user *models.User) error {
	userDocument := mappers.ToUserDocument(user)

	userDocument.Version = 1

	if _, err := u.userCollection.InsertOne(ctx, userDocument); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
