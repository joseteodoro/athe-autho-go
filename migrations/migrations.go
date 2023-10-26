package migrations

import (
	"athe-autho-go/models"

	"github.com/jinzhu/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(models.Realm{}, models.User{}, models.Role{}, models.UserGroup{}, models.UserGroupMembership{}, models.UserRoleMembership{})
}
