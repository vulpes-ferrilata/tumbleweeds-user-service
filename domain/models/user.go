package models

import (
	"github.com/VulpesFerrilata/user-service/domain/models/common"
	"github.com/google/uuid"
)

func NewUser(id uuid.UUID, displayName string, version int) *User {
	user := new(User)
	user.Entity = common.NewEntity(id, version)
	user.displayName = displayName
	return user
}

type User struct {
	common.Entity
	displayName string
}

func (u User) GetDisplayName() string {
	return u.displayName
}
