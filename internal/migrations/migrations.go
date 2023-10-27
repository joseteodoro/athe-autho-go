package migrations

import (
	db "athe-autho-go/internal/database"
	"athe-autho-go/internal/models"
)

func Migrate() {
	connection := db.Init()
	connection.AutoMigrate(models.Realm{}, models.User{}, models.Attribute{}, models.UserGroup{}, models.UserGroupMembership{}, models.UserAttributeMembership{})
}
