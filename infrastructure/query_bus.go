package infrastructure

import (
	"github.com/vulpes-ferrilata/cqrs"
	"github.com/vulpes-ferrilata/user-service/application/queries"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/middlewares"
)

func NewQueryBus(validationMiddleware *middlewares.ValidationMiddleware,
	getUserByIDQueryHandler *queries.GetUserByIDQueryHandler) (*cqrs.QueryBus, error) {
	queryBus := &cqrs.QueryBus{}

	queryBus.Use(
		validationMiddleware.QueryHandlerMiddleware(),
	)

	queryBus.Register(&queries.GetUserByIDQuery{}, cqrs.WrapQueryHandlerFunc(getUserByIDQueryHandler.Handle))

	return queryBus, nil
}
