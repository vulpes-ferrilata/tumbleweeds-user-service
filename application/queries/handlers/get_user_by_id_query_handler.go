package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/user-service/application/queries"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/query/wrappers"
	"github.com/vulpes-ferrilata/user-service/view/models"
	"github.com/vulpes-ferrilata/user-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewGetUserByIDQueryHandler(validate *validator.Validate, userProjector projectors.UserProjector) query.QueryHandler[*queries.GetUserByID, *models.User] {
	handler := &getUserByIDQueryHandler{
		userProjector: userProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetUserByID, *models.User](validate, handler)

	return validationWrapper
}

type getUserByIDQueryHandler struct {
	userProjector projectors.UserProjector
}

func (g getUserByIDQueryHandler) Handle(ctx context.Context, getUserByIDQuery *queries.GetUserByID) (*models.User, error) {
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
