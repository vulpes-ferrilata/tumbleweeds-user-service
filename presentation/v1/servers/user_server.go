package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/shared/proto/v1/user"
	"github.com/vulpes-ferrilata/user-service/application/commands"
	"github.com/vulpes-ferrilata/user-service/application/queries"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/user-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/user-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewUserServer(getUserByIDQueryHandler query.QueryHandler[*queries.GetUserByIDQuery, *models.User],
	createUserCommandHandler command.CommandHandler[*commands.CreateUserCommand]) user.UserServer {
	return &userServer{
		getUserByIDQueryHandler:  getUserByIDQueryHandler,
		createUserCommandHandler: createUserCommandHandler,
	}
}

type userServer struct {
	user.UnimplementedUserServer
	getUserByIDQueryHandler  query.QueryHandler[*queries.GetUserByIDQuery, *models.User]
	createUserCommandHandler command.CommandHandler[*commands.CreateUserCommand]
}

func (u userServer) GetUserByID(ctx context.Context, getUserByUserIDRequest *user.GetUserByIDRequest) (*user.UserResponse, error) {
	getUserByIDQuery := &queries.GetUserByIDQuery{
		ID: getUserByUserIDRequest.GetID(),
	}

	user, err := u.getUserByIDQueryHandler.Handle(ctx, getUserByIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userResponse := mappers.ToUserResponse(user)

	return userResponse, nil
}

func (u userServer) CreateUser(ctx context.Context, createUserRequest *user.CreateUserRequest) (*emptypb.Empty, error) {
	createUserCommand := &commands.CreateUserCommand{
		ID:          createUserRequest.GetID(),
		DisplayName: createUserRequest.GetDisplayName(),
	}

	if err := u.createUserCommandHandler.Handle(ctx, createUserCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
