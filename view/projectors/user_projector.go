package projectors

import (
	"context"

	"github.com/VulpesFerrilata/user-service/persistence/entities"
	"github.com/VulpesFerrilata/user-service/persistence/repositories"
	"github.com/VulpesFerrilata/user-service/view/mappers"
	"github.com/VulpesFerrilata/user-service/view/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserProjector interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

func NewUserProjector(userRepository repositories.Repository[entities.User],
	userMapper mappers.UserMapper) UserProjector {
	return &userProjector{
		userRepository: userRepository,
		userMapper:     userMapper,
	}
}

type userProjector struct {
	userRepository repositories.Repository[entities.User]
	userMapper     mappers.UserMapper
}

func (u userProjector) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	userEntity, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := u.userMapper.ToModel(ctx, userEntity)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
