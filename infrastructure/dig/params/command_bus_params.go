package params

import (
	"github.com/VulpesFerrilata/user-service/infrastructure/bus"
	"go.uber.org/dig"
)

type CommandBusParams struct {
	dig.In

	CommandHandlers []bus.CommandHandler `group:"commandBus"`
}
