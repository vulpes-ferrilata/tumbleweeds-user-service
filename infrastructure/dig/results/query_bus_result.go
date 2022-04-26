package results

import (
	"github.com/VulpesFerrilata/user-service/infrastructure/bus"
	"go.uber.org/dig"
)

type QueryHandlerResult struct {
	dig.Out

	QueryHandler bus.QueryHandler `group:"queryBus"`
}
