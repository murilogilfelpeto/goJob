package configuration

import (
	"github.com/murilogilfelpeto/goJob/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySql() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	logger.Info("Initializing MySQL connection")
	const url = "root:admin@tcp(localhost:3306)/go-job?parseTime=true"
	connection, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing MySQL connection: %v", err)
		return nil, err
	}
	logger.Info("MySQL connection initialized")
	logger.Info("Running Migrations")
	err = connection.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error running migrations: %v", err)
		return nil, err
	}
	logger.Info("Migrations ran successfully")
	return connection, nil
}
