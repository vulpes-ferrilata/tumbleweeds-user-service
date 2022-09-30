package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/user-service-proto/pb"
	"github.com/vulpes-ferrilata/user-service-proto/pb/requests"
	"github.com/vulpes-ferrilata/user-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/user-service/application/commands"
	"github.com/vulpes-ferrilata/user-service/application/queries"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/user-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/user-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewUserServer(getUserByIDQueryHandler query.QueryHandler[*queries.GetUserByID, *models.User],
	createUserCommandHandler command.CommandHandler[*commands.CreateUser]) pb.UserServer {
	return &userServer{
		getUserByIDQueryHandler:  getUserByIDQueryHandler,
		createUserCommandHandler: createUserCommandHandler,
	}
}

type userServer struct {
	pb.UnimplementedUserServer
	getUserByIDQueryHandler  query.QueryHandler[*queries.GetUserByID, *models.User]
	createUserCommandHandler command.CommandHandler[*commands.CreateUser]
}

func (u userServer) GetUserByID(ctx context.Context, getUserByUserIDRequest *requests.GetUserByID) (*responses.User, error) {
	getUserByIDQuery := &queries.GetUserByID{
		UserID: getUserByUserIDRequest.GetUserID(),
	}

	user, err := u.getUserByIDQueryHandler.Handle(ctx, getUserByIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userResponse := mappers.ToUserResponse(user)

	return userResponse, nil
}

func (u userServer) CreateUser(ctx context.Context, createUserRequest *requests.CreateUser) (*emptypb.Empty, error) {
	createUserCommand := &commands.CreateUser{
		UserID:      createUserRequest.GetUserID(),
		DisplayName: createUserRequest.GetDisplayName(),
	}

	if err := u.createUserCommandHandler.Handle(ctx, createUserCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
