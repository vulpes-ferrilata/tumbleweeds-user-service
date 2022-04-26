package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(config *Config) (*gorm.DB, error) {
	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp" + "(" + config.Database.Host + ":" + config.Database.Port + ")" + "/" + config.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialector := mysql.Open(dsn)
	return gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
