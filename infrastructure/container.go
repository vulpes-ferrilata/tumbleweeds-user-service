package infrastructure

import (
	command_handlers "github.com/VulpesFerrilata/user-service/application/commands/handlers"
	query_handlers "github.com/VulpesFerrilata/user-service/application/queries/handlers"
	"github.com/VulpesFerrilata/user-service/domain/aggregators"

	domain_mappers "github.com/VulpesFerrilata/user-service/domain/mappers"
	domain_services "github.com/VulpesFerrilata/user-service/domain/services"
	"github.com/VulpesFerrilata/user-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/user-service/infrastructure/persistence/repositories"
	"github.com/VulpesFerrilata/user-service/persistence/entities"
	view_mappers "github.com/VulpesFerrilata/user-service/view/mappers"
	"github.com/VulpesFerrilata/user-service/view/projectors"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//3rd party libraries
	container.Provide(NewConfig)
	container.Provide(NewGorm)
	container.Provide(NewValidate)
	container.Provide(NewCommandBus)
	container.Provide(NewQueryBus)
	container.Provide(NewUniversalTranslator)

	//middlewares
	container.Provide(middlewares.NewTransactionMiddleware)
	container.Provide(middlewares.NewValidationMiddleware)
	container.Provide(middlewares.NewErrorHandlerMiddleware)
	container.Provide(middlewares.NewTranslatorMiddleware)

	//Persistence layer
	container.Provide(repositories.NewRepository[entities.User])

	//View layer
	//--Projectors
	container.Provide(projectors.NewUserProjector)
	//--Mappers
	container.Provide(view_mappers.NewUserMapper)

	//Domain layer
	//--Services
	container.Provide(domain_services.NewUserService)
	//--Aggregators
	container.Provide(aggregators.NewUserAggregator)
	//--Mappers
	container.Provide(domain_mappers.NewUserMapper)

	//Application layer
	//--Queries
	container.Provide(query_handlers.NewGetUserQueryHandler)
	//--Commands
	container.Provide(command_handlers.NewCreateUserCommandHandler)

	return container
}
