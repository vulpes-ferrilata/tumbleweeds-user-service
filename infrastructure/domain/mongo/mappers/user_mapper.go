package mappers

import (
	"github.com/vulpes-ferrilata/user-service/domain/models"
	"github.com/vulpes-ferrilata/user-service/infrastructure/domain/mongo/documents"
)

func ToUserDocument(user *models.User) *documents.User {
	if user == nil {
		return nil
	}

	return &documents.User{
		DocumentRoot: documents.DocumentRoot{
			Document: documents.Document{
				ID: user.GetID(),
			},
			Version: user.GetVersion(),
		},
		DisplayName: user.GetDisplayName(),
	}
}

func ToUserDomain(userDocument *documents.User) *models.User {
	if userDocument == nil {
		return nil
	}

	return models.UserBuilder{}.
		SetID(userDocument.ID).
		SetDisplayName(userDocument.DisplayName).
		SetVersion(userDocument.Version).
		Create()
}
