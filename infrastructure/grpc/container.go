package grpc

import (
	"github.com/VulpesFerrilata/user-service/infrastructure"
	grpc_handlers "github.com/VulpesFerrilata/user-service/presentation/grpc/handlers"
	"github.com/VulpesFerrilata/user-service/presentation/grpc/mappers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := infrastructure.NewContainer()

	//Presentation layer
	//--Handlers
	container.Provide(grpc_handlers.NewUserHandler)
	//--Mappers
	container.Provide(mappers.NewUserMapper)

	container.Provide(NewServer)
	container.Provide(NewClient)
	container.Provide(NewService)

	return container
}
