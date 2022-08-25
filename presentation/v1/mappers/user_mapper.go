package mappers

import (
	"github.com/vulpes-ferrilata/shared/proto/v1/user"
	"github.com/vulpes-ferrilata/user-service/view/models"
)

func ToUserResponse(userView *models.User) *user.UserResponse {
	if userView == nil {
		return nil
	}

	return &user.UserResponse{
		ID:          userView.ID.Hex(),
		DisplayName: userView.DisplayName,
	}
}
