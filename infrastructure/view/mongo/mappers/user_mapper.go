package mappers

import (
	"github.com/vulpes-ferrilata/user-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/user-service/view/models"
)

func ToUserView(userDocument *documents.User) *models.User {
	if userDocument == nil {
		return nil
	}

	user := &models.User{
		ID:          userDocument.ID,
		DisplayName: userDocument.DisplayName,
	}

	return user
}
