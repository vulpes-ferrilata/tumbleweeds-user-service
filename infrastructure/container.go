package infrastructure

import (
	command_handlers "github.com/vulpes-ferrilata/user-service/application/commands/handlers"
	query_handlers "github.com/vulpes-ferrilata/user-service/application/queries/handlers"
	"github.com/vulpes-ferrilata/user-service/infrastructure/domain/mongo/repositories"
	"github.com/vulpes-ferrilata/user-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/user-service/infrastructure/view/mongo/projectors"
	"github.com/vulpes-ferrilata/user-service/presentation"
	"github.com/vulpes-ferrilata/user-service/presentation/v1/servers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//Infrastructure layer
	container.Provide(NewConfig)
	container.Provide(NewMongo)
	container.Provide(NewValidator)
	container.Provide(NewLogrus)
	container.Provide(NewUniversalTranslator)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)

	//Domain layer
	//--Repositories
	container.Provide(repositories.NewUserRepository)

	//View layer
	//--Projectors
	container.Provide(projectors.NewUserProjector)

	//Application layer
	//--Queries
	container.Provide(query_handlers.NewGetUserByIDQueryHandler)
	//--Commands
	container.Provide(command_handlers.NewCreateUserCommandHandler)

	//Presentation layer
	//--Server
	container.Provide(presentation.NewServer)
	//--Controllers
	container.Provide(servers.NewUserServer)

	return container
}
