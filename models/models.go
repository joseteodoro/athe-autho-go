package models

import (
	"github.com/jinzhu/gorm"
)

type NamedEntity struct {
	Name        string `gorm:"not null"`
	Description string
}

type BaseEntity struct {
	gorm.Model
	Uuid    string `gorm:"not null"`
	RealmID uint   `gorm:"not null"`
}

type Realm struct {
	gorm.Model
	Uuid string `gorm:"not null"`
	NamedEntity
}

type User struct {
	BaseEntity
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type UserGroup struct {
	BaseEntity
	NamedEntity
}

type UserGroupMembership struct {
	BaseEntity
	UserID      uint `gorm:"not null"`
	UserGroupID uint `gorm:"not null"`
}

type UserRoleMembership struct {
	BaseEntity
	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
}

type Role struct {
	BaseEntity
	NamedEntity
}
