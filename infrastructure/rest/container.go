package rest

import (
	"github.com/VulpesFerrilata/user-service/infrastructure"
	"github.com/VulpesFerrilata/user-service/presentation/rest/controllers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := infrastructure.NewContainer()

	//Presentation layer
	container.Provide(controllers.NewUserController)

	container.Provide(NewRouter)
	container.Provide(NewApp)
	container.Provide(NewServer)
	container.Provide(NewClient)
	container.Provide(NewService)

	return container
}
