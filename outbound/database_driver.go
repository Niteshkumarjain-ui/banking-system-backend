package outbound

import (
	"banking-system-backend/util"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DatabaseDriver *gorm.DB

// function to create database client
func createDatabaseClient() {
	logger := util.GetLogger()

	var err error

	var db *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", util.Configuration.Database.Host, util.Configuration.Database.Username, util.Configuration.Database.Password, util.Configuration.Database.DBName, util.Configuration.Database.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", util.Configuration.Database.Schema), // schema name
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Errorf("Database connection error %v", err)
	} else {
		logger.Infof("Database connection successful")
	}

	DatabaseDriver = db
}

func init() {
	createDatabaseClient()
}
