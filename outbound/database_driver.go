package outbound

import (
	"banking-system-backend/util"
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
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

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		logger.Errorf("Database tracing setup error %v", err)
	}

	DatabaseDriver = db
}

func init() {
	createDatabaseClient()
}
