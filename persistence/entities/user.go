package entities

import (
	"github.com/VulpesFerrilata/user-service/persistence/entities/common"
)

type User struct {
	common.Entity
	DisplayName string `gorm:"type:varchar(20); unique"`
}
