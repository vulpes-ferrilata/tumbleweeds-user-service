package infrastructure

import (
	"github.com/vulpes-ferrilata/user-service/application/commands"
	"github.com/vulpes-ferrilata/user-service/application/queries"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/user-service/infrastructure/domain/mongo/repositories"
	"github.com/vulpes-ferrilata/user-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/user-service/infrastructure/view/mongo/projectors"
	"github.com/vulpes-ferrilata/user-service/presentation"
	v1 "github.com/vulpes-ferrilata/user-service/presentation/v1"
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
	container.Provide(NewCommandBus)
	container.Provide(NewQueryBus)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)
	//--Cqrs middlewares
	container.Provide(middlewares.NewValidationMiddleware)
	container.Provide(middlewares.NewTransactionMiddleware)

	//Domain layer
	//--Repositories
	container.Provide(repositories.NewUserRepository)

	//View layer
	//--Projectors
	container.Provide(projectors.NewUserProjector)

	//Application layer
	//--Queries
	container.Provide(queries.NewGetUserByIDQueryHandler)
	//--Commands
	container.Provide(commands.NewCreateUserCommandHandler)

	//Presentation layer
	container.Provide(presentation.NewServer)
	container.Provide(v1.NewUserServer)

	return container
}
