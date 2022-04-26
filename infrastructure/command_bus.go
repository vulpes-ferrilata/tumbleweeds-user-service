package infrastructure

import (
	"github.com/VulpesFerrilata/user-service/infrastructure/bus"
	"github.com/VulpesFerrilata/user-service/infrastructure/dig/params"
	"github.com/VulpesFerrilata/user-service/infrastructure/middlewares"
	"github.com/pkg/errors"
)

func NewCommandBus(params params.CommandBusParams,
	validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware) (bus.CommandBus, error) {
	commandBus := bus.NewCommandBus()

	if err := commandBus.Register(params.CommandHandlers...); err != nil {
		return nil, errors.WithStack(err)
	}

	commandBus.Use(
		transactionMiddleware.WrapCommandHandler,
		validationMiddleware.WrapCommandHandler,
	)

	return commandBus, nil
}
