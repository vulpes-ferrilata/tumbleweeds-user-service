package params

import (
	"github.com/VulpesFerrilata/user-service/infrastructure/bus"
	"go.uber.org/dig"
)

type QueryBusParams struct {
	dig.In

	QueryHandlers []bus.QueryHandler `group:"queryBus"`
}
