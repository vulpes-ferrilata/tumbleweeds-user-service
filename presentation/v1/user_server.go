package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/cqrs"
	"github.com/vulpes-ferrilata/user-service-proto/pb"
	pb_models "github.com/vulpes-ferrilata/user-service-proto/pb/models"
	"github.com/vulpes-ferrilata/user-service/application/commands"
	"github.com/vulpes-ferrilata/user-service/application/queries"
	"github.com/vulpes-ferrilata/user-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/user-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewUserServer(commandBus *cqrs.CommandBus,
	queryBus *cqrs.QueryBus) pb.UserServer {
	return &userServer{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

type userServer struct {
	pb.UnimplementedUserServer
	queryBus   *cqrs.QueryBus
	commandBus *cqrs.CommandBus
}

func (u userServer) GetUserByID(ctx context.Context, getUserByUserIDRequest *pb_models.GetUserByID) (*pb_models.User, error) {
	getUserByIDQuery := &queries.GetUserByIDQuery{
		UserID: getUserByUserIDRequest.GetUserID(),
	}

	user, err := cqrs.ParseQueryHandlerFunc[*queries.GetUserByIDQuery, *models.User](u.queryBus.Execute)(ctx, getUserByIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userResponse, err := mappers.UserMapper{}.ToResponse(user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return userResponse, nil
}

func (u userServer) CreateUser(ctx context.Context, createUserRequest *pb_models.CreateUser) (*emptypb.Empty, error) {
	createUserCommand := &commands.CreateUserCommand{
		UserID:      createUserRequest.GetUserID(),
		DisplayName: createUserRequest.GetDisplayName(),
	}

	if err := u.commandBus.Execute(ctx, createUserCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
