package models

import (
	"github.com/jinzhu/gorm"
)

type NamedEntity struct {
	Name        string `gorm:"not null"`
	Description string
}

type ReadWriteEntity struct {
	CanRead  bool `gorm:"not null; default:false"`
	CanWrite bool `gorm:"not null; default:false"`
	CanAdmin bool `gorm:"not null; default:false"`
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
	Username   string      `gorm:"not null"`
	Password   string      `gorm:"not null"`
	Groups     []UserGroup `gorm:"many2many:user_group_membership;"`
	Attributes []Attribute `gorm:"many2many:user_attribute_membership;"`
}

type UserGroup struct {
	BaseEntity
	NamedEntity
}

type UserGroupMembership struct {
	BaseEntity
	ReadWriteEntity
	UserID      uint `gorm:"not null"`
	UserGroupID uint `gorm:"not null"`
}

type UserAttributeMembership struct {
	BaseEntity
	ReadWriteEntity
	UserID      uint `gorm:"not null"`
	AttributeID uint `gorm:"not null"`
}

type Attribute struct {
	BaseEntity
	NamedEntity
}
