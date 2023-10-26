package db

import (
	"athe-autho-go/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() *gorm.DB {
	config := configs.NewDatabaseConfig()
	db, err := gorm.Open(config.Dialect, config.ConnectionString)
	if err != nil {
		panic(err)
	}

	return db
}
