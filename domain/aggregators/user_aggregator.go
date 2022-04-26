package aggregators

import (
	"context"

	"github.com/VulpesFerrilata/user-service/domain/mappers"
	"github.com/VulpesFerrilata/user-service/domain/models"
	"github.com/VulpesFerrilata/user-service/persistence/entities"
	"github.com/VulpesFerrilata/user-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserAggregator interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	Save(ctx context.Context, user *models.User) error
}

func NewUserAggregator(userRepository repositories.Repository[entities.User],
	userMapper mappers.UserMapper) UserAggregator {
	return &userAggregator{
		userRepository: userRepository,
		userMapper:     userMapper,
	}
}

type userAggregator struct {
	userRepository repositories.Repository[entities.User]
	userMapper     mappers.UserMapper
}

func (u userAggregator) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
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

func (u userAggregator) Save(ctx context.Context, user *models.User) error {
	userEntity, err := u.userMapper.ToEntity(ctx, user)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := u.userRepository.Save(ctx, userEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
