package handlers

import (
	"context"

	"github.com/VulpesFerrilata/user-service/application/queries"
	"github.com/VulpesFerrilata/user-service/infrastructure/dig/results"
	"github.com/VulpesFerrilata/user-service/view/projectors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewGetUserQueryHandler(userProjector projectors.UserProjector) results.QueryHandlerResult {
	queryHandler := &getUserByIDQueryHandler{
		userProjector: userProjector,
	}

	return results.QueryHandlerResult{
		QueryHandler: queryHandler,
	}
}

type getUserByIDQueryHandler struct {
	userProjector projectors.UserProjector
}

func (g getUserByIDQueryHandler) GetQuery() interface{} {
	return &queries.GetUserByIDQuery{}
}

func (g getUserByIDQueryHandler) Handle(ctx context.Context, query interface{}) (interface{}, error) {
	getUserByIDQuery := query.(queries.GetUserByIDQuery)

	userID, err := uuid.Parse(getUserByIDQuery.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := g.userProjector.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
