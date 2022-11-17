package mappers

import (
	"github.com/vulpes-ferrilata/user-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/user-service/view/models"
)

type UserMapper struct{}

func (u UserMapper) ToResponse(userView *models.User) (*responses.User, error) {
	if userView == nil {
		return nil, nil
	}

	return &responses.User{
		ID:          userView.ID.Hex(),
		DisplayName: userView.DisplayName,
	}, nil
}
