package mappers

import (
	pb_models "github.com/vulpes-ferrilata/user-service-proto/pb/models"
	"github.com/vulpes-ferrilata/user-service/view/models"
)

type UserMapper struct{}

func (u UserMapper) ToResponse(userView *models.User) (*pb_models.User, error) {
	if userView == nil {
		return nil, nil
	}

	return &pb_models.User{
		ID:          userView.ID.Hex(),
		DisplayName: userView.DisplayName,
	}, nil
}
