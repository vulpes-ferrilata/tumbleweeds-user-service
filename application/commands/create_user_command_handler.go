package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/user-service/domain/models"
	"github.com/vulpes-ferrilata/user-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUser struct {
	UserID      string `validate:"required,objectid"`
	DisplayName string `validate:"required,min=1,max=20"`
}

func NewCreateUserCommandHandler(userRepository repositories.UserRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		userRepository: userRepository,
	}
}

type CreateUserCommandHandler struct {
	userRepository repositories.UserRepository
}

func (c CreateUserCommandHandler) Handle(ctx context.Context, createUserCommand *CreateUser) error {
	userID, err := primitive.ObjectIDFromHex(createUserCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	user := models.UserBuilder{}.
		SetID(userID).
		SetDisplayName(createUserCommand.DisplayName).
		Create()

	if err := c.userRepository.Insert(ctx, user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
