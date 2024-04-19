// db.go

package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig holds the configuration details for the database connection
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// NewDBConnection creates a new database connection based on the provided configuration
func NewDBConnection(config *DBConfig) (*gorm.DB, error) {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
