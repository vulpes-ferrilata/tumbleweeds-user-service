package rest

import (
	"github.com/VulpesFerrilata/go-micro/plugins/client/grpc/v3"
	"github.com/VulpesFerrilata/user-service/infrastructure/middlewares"
	"github.com/asim/go-micro/v3/client"
)

func NewClient(translatorMiddleware *middlewares.TranslatorMiddleware, errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware) client.Client {
	client := grpc.NewClient(
		client.WrapCall(translatorMiddleware.WrapCall),
		client.WrapCall(errorHandlerMiddleware.WrapCall),
	)

	return client
}
