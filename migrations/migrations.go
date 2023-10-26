package migrations

import (
	"athe-autho-go/models"

	"github.com/jinzhu/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(models.User{})
}
