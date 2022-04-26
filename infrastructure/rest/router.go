package rest

import (
	"github.com/VulpesFerrilata/user-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/user-service/presentation/rest/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Router interface {
	Init(app *iris.Application)
}

func NewRouter(translatorMiddleware *middlewares.TranslatorMiddleware,
	errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware,
	userController controllers.UserController) Router {
	return &router{
		translatorMiddleware:   translatorMiddleware,
		errorHandlerMiddleware: errorHandlerMiddleware,
		userController:         userController,
	}
}

type router struct {
	translatorMiddleware   *middlewares.TranslatorMiddleware
	errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware
	userController         controllers.UserController
}

func (r router) Init(app *iris.Application) {
	api := app.Party("/api")

	userApi := api.Party("/user")
	userApi.Use(r.translatorMiddleware.Serve)
	userMvc := mvc.New(userApi)
	userMvc.HandleError(r.errorHandlerMiddleware.Handle)
	userMvc.Handle(r.userController)
}
