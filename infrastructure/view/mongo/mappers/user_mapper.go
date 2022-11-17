package mappers

import (
	"github.com/vulpes-ferrilata/user-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/user-service/view/models"
)

type UserMapper struct{}

func (u UserMapper) ToView(userDocument *documents.User) (*models.User, error) {
	if userDocument == nil {
		return nil, nil
	}

	return &models.User{
		ID:          userDocument.ID,
		DisplayName: userDocument.DisplayName,
	}, nil
}
