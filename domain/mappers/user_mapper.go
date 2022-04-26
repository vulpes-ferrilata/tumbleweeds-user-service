package mappers

import (
	"context"

	"github.com/VulpesFerrilata/user-service/domain/models"
	"github.com/VulpesFerrilata/user-service/infrastructure/gorm"
	"github.com/VulpesFerrilata/user-service/persistence/entities"
	"github.com/VulpesFerrilata/user-service/persistence/entities/common"
)

type UserMapper interface {
	ToEntity(ctx context.Context, user *models.User) (*entities.User, error)
	ToModel(ctx context.Context, userEntity *entities.User) (*models.User, error)
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

type userMapper struct{}

func (u *userMapper) ToEntity(ctx context.Context, user *models.User) (*entities.User, error) {
	if user == nil {
		return nil, nil
	}

	userEntity := &entities.User{
		Entity: common.Entity{
			ID:      user.GetID(),
			Version: gorm.Version(user.GetVersion()),
		},
		DisplayName: user.GetDisplayName(),
	}

	return userEntity, nil
}

func (u *userMapper) ToModel(ctx context.Context, userEntity *entities.User) (*models.User, error) {
	if userEntity == nil {
		return nil, nil
	}

	user := models.NewUser(
		userEntity.ID,
		userEntity.DisplayName,
		int(userEntity.Version),
	)

	return user, nil
}
