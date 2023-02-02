package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/user-service/view/models"
	"github.com/vulpes-ferrilata/user-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserByIDQuery struct {
	UserID string `validate:"required,objectid"`
}

func NewGetUserByIDQueryHandler(userProjector projectors.UserProjector) *GetUserByIDQueryHandler {
	return &GetUserByIDQueryHandler{
		userProjector: userProjector,
	}
}

type GetUserByIDQueryHandler struct {
	userProjector projectors.UserProjector
}

func (g GetUserByIDQueryHandler) Handle(ctx context.Context, getUserByIDQuery *GetUserByIDQuery) (*models.User, error) {
	userID, err := primitive.ObjectIDFromHex(getUserByIDQuery.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := g.userProjector.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
