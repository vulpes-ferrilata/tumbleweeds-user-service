package mappers

import (
	"context"

	"github.com/VulpesFerrilata/shared/proto/user"
	"github.com/VulpesFerrilata/user-service/view/models"
)

type UserMapper interface {
	ToResponse(ctx context.Context, userView *models.User) (*user.UserResponse, error)
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

type userMapper struct{}

func (u userMapper) ToResponse(ctx context.Context, userView *models.User) (*user.UserResponse, error) {
	if userView == nil {
		return nil, nil
	}

	userResponse := &user.UserResponse{
		ID:          userView.ID,
		DisplayName: userView.DisplayName,
	}

	return userResponse, nil
}
