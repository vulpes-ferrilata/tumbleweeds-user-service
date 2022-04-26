package handlers

import (
	"context"

	"github.com/VulpesFerrilata/user-service/application/commands"
	"github.com/VulpesFerrilata/user-service/domain/aggregators"
	domain_services "github.com/VulpesFerrilata/user-service/domain/services"
	"github.com/VulpesFerrilata/user-service/infrastructure/dig/results"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewCreateUserCommandHandler(userAggregator aggregators.UserAggregator,
	userService domain_services.UserService) results.CommandHandlerResult {
	commandHandler := &createUserCommandHandler{
		userAggregator: userAggregator,
		userService:    userService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type createUserCommandHandler struct {
	userAggregator aggregators.UserAggregator
	userService    domain_services.UserService
}

func (c createUserCommandHandler) GetCommand() interface{} {
	return &commands.CreateUserCommand{}
}

func (c createUserCommandHandler) Handle(ctx context.Context, command interface{}) error {
	createUserCommand := command.(*commands.CreateUserCommand)

	userID, err := uuid.Parse(createUserCommand.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	user, err := c.userService.NewUser(ctx, userID, createUserCommand.DisplayName)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.userAggregator.Save(ctx, user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
