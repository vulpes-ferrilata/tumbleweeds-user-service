package mappers

import (
	"context"

	"github.com/VulpesFerrilata/user-service/persistence/entities"
	"github.com/VulpesFerrilata/user-service/view/models"
)

type UserMapper interface {
	ToModel(ctx context.Context, userEntity *entities.User) (*models.User, error)
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

type userMapper struct{}

func (u userMapper) ToModel(ctx context.Context, userEntity *entities.User) (*models.User, error) {
	if userEntity == nil {
		return nil, nil
	}

	user := &models.User{
		ID:          userEntity.ID.String(),
		DisplayName: userEntity.DisplayName,
	}

	return user, nil
}
