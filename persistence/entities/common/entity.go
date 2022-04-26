package common

import (
	"github.com/VulpesFerrilata/user-service/infrastructure/gorm"
	"github.com/google/uuid"
)

type Entity struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Version gorm.Version
}
