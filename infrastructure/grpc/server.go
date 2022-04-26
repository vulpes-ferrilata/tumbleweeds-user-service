package grpc

import (
	"github.com/VulpesFerrilata/go-micro/plugins/server/grpc/v3"
	"github.com/VulpesFerrilata/shared/proto/user"
	"github.com/VulpesFerrilata/user-service/infrastructure/middlewares"
	"github.com/asim/go-micro/v3/server"
	"github.com/pkg/errors"
)

func NewServer(translatorMiddleware *middlewares.TranslatorMiddleware,
	errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware, userHandler user.UserHandler) (server.Server, error) {
	server := grpc.NewServer(
		server.WrapHandler(translatorMiddleware.WrapHandler),
		server.WrapHandler(errorHandlerMiddleware.WrapHandler),
	)

	if err := user.RegisterUserHandler(server, userHandler); err != nil {
		return nil, errors.WithStack(err)
	}

	return server, nil
}
