package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/user-service/application/commands"
	"github.com/vulpes-ferrilata/user-service/domain/models"
	"github.com/vulpes-ferrilata/user-service/domain/repositories"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/user-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCreateUserCommandHandler(validate *validator.Validate, db *mongo.Database, userRepository repositories.UserRepository) command.CommandHandler[*commands.CreateUserCommand] {
	handler := &createUserCommandHandler{
		userRepository: userRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.CreateUserCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type createUserCommandHandler struct {
	userRepository repositories.UserRepository
}

func (c createUserCommandHandler) GetCommand() interface{} {
	return &commands.CreateUserCommand{}
}

func (c createUserCommandHandler) Handle(ctx context.Context, createUserCommand *commands.CreateUserCommand) error {
	id, err := primitive.ObjectIDFromHex(createUserCommand.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	user := models.UserBuilder{}.
		SetID(id).
		SetDisplayName(createUserCommand.DisplayName).
		Create()

	if err := c.userRepository.Insert(ctx, user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
