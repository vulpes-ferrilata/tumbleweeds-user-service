package validators

import (
	"context"

	"github.com/VulpesFerrilata/user-service/application/commands"
)

func NewCreateUserCommandValidator() Validator[commands.CreateUserCommand] {
	return &createUserCommandValidator{}
}

type createUserCommandValidator struct {
}

func (c createUserCommandValidator) Validate(ctx context.Context, createUserCommand commands.CreateUserCommand) error {

}
