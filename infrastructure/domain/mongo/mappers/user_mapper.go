package mappers

import (
	"github.com/vulpes-ferrilata/user-service/domain/models"
	"github.com/vulpes-ferrilata/user-service/infrastructure/domain/mongo/documents"
)

type UserMapper struct{}

func (u UserMapper) ToDocument(user *models.User) (*documents.User, error) {
	if user == nil {
		return nil, nil
	}

	return &documents.User{
		DocumentRoot: documents.DocumentRoot{
			Document: documents.Document{
				ID: user.GetID(),
			},
			Version: user.GetVersion(),
		},
		DisplayName: user.GetDisplayName(),
	}, nil
}
