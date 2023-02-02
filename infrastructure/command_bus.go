package infrastructure

import (
	"github.com/vulpes-ferrilata/cqrs"
	"github.com/vulpes-ferrilata/user-service/application/commands"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/middlewares"
)

func NewCommandBus(validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware,
	createUserCommandHandler *commands.CreateUserCommandHandler) (*cqrs.CommandBus, error) {
	commandBus := &cqrs.CommandBus{}

	commandBus.Use(
		validationMiddleware.CommandHandlerMiddleware(),
		transactionMiddleware.CommandHandlerMiddleware(),
	)

	commandBus.Register(&commands.CreateUser{}, cqrs.WrapCommandHandlerFunc(createUserCommandHandler.Handle))

	return commandBus, nil
}
