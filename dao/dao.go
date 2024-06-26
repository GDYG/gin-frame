// dao.go

package dao

import (
	"gin-frame/config"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func InitDB() *gorm.DB {
	dbConfig := config.GetDBConfig()

	DB, Err = config.NewDBConnection(dbConfig)
	if Err != nil {
		panic("Failed to connect database")
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
