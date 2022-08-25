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

func NewGetUserByIDQueryHandler(validate *validator.Validate, userProjector projectors.UserProjector) query.QueryHandler[*queries.GetUserByIDQuery, *models.User] {
	handler := &getUserByIDQueryHandler{
		userProjector: userProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetUserByIDQuery, *models.User](validate, handler)

	return validationWrapper
}

type getUserByIDQueryHandler struct {
	userProjector projectors.UserProjector
}

func (g getUserByIDQueryHandler) Handle(ctx context.Context, getUserByIDQuery *queries.GetUserByIDQuery) (*models.User, error) {
	id, err := primitive.ObjectIDFromHex(getUserByIDQuery.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := g.userProjector.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
