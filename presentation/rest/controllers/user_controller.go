package controllers

import (
	"github.com/VulpesFerrilata/user-service/application/commands"
	"github.com/VulpesFerrilata/user-service/infrastructure/bus"
	"github.com/VulpesFerrilata/user-service/presentation/rest/requests"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/pkg/errors"
)

type UserController interface {
	Post(ctx iris.Context) (mvc.Result, error)
	PostBy(ctx iris.Context, id string) (mvc.Result, error)
}

func NewUserController(commandBus bus.CommandBus) UserController {
	return &userController{
		commandBus: commandBus,
	}
}

type userController struct {
	commandBus bus.CommandBus
}

func (u userController) Post(ctx iris.Context) (mvc.Result, error) {
	userRegisterRequest := &requests.UserRequest{}

	if err := ctx.ReadJSON(userRegisterRequest); err != nil {
		return nil, errors.WithStack(err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	createUserCommand := &commands.CreateUserCommand{
		ID:          id.String(),
		DisplayName: userRegisterRequest.DisplayName,
	}

	if err := u.commandBus.Execute(ctx.Request().Context(), createUserCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return mvc.Response{
		Code:   iris.StatusCreated,
		Object: id,
	}, nil
}

func (u userController) PostBy(ctx iris.Context, id string) (mvc.Result, error) {
	return mvc.Response{
		Code: iris.StatusNotImplemented,
	}, nil
}
