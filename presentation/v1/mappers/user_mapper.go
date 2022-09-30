package mappers

import (
	"github.com/vulpes-ferrilata/user-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/user-service/view/models"
)

func ToUserResponse(userView *models.User) *responses.User {
	if userView == nil {
		return nil
	}

	return &responses.User{
		ID:          userView.ID.Hex(),
		DisplayName: userView.DisplayName,
	}
}
