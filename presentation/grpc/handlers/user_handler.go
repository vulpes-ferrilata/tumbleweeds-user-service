package handlers

import (
	"context"

	"github.com/VulpesFerrilata/shared/proto/user"
	"github.com/VulpesFerrilata/user-service/application/commands"
	"github.com/VulpesFerrilata/user-service/application/queries"
	"github.com/VulpesFerrilata/user-service/infrastructure/bus"
	"github.com/VulpesFerrilata/user-service/presentation/grpc/mappers"
	"github.com/VulpesFerrilata/user-service/view/models"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

func NewUserHandler(queryBus bus.QueryBus,
	commandBus bus.CommandBus,
	userMapper mappers.UserMapper) user.UserHandler {
	return &userHandler{
		queryBus:   queryBus,
		commandBus: commandBus,
		userMapper: userMapper,
	}
}

type userHandler struct {
	queryBus   bus.QueryBus
	commandBus bus.CommandBus
	userMapper mappers.UserMapper
}

func (u userHandler) GetUserByID(ctx context.Context, request *user.GetUserByIDRequest, response *user.UserResponse) error {
	getUserByIDQuery := &queries.GetUserByIDQuery{
		ID: request.GetID(),
	}

	result, err := u.queryBus.Execute(ctx, getUserByIDQuery)
	if err != nil {
		return errors.WithStack(err)
	}
	userView := result.(*models.User)

	userResponse, err := u.userMapper.ToResponse(ctx, userView)
	if err != nil {
		return errors.WithStack(err)
	}
	*response = *userResponse

	return nil
}

func (u userHandler) CreateUser(ctx context.Context, createUserRequest *user.CreateUserRequest, _ *empty.Empty) error {
	createUserCommand := &commands.CreateUserCommand{
		ID:          createUserRequest.GetID(),
		DisplayName: createUserRequest.GetDisplayName(),
	}
	if err := u.commandBus.Execute(ctx, createUserCommand); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
