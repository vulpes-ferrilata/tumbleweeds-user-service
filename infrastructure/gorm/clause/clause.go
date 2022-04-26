package clause

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClauseInterface interface {
	clause.Interface
	gorm.StatementModifier
}
