package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConnection struct {
	Database *gorm.DB
}

func NewDatabaseConnection(envData *Env) *DatabaseConnection {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kathmandu", envData.PostgresHost, envData.PostgresUser, envData.PostgresPassword, envData.PostgresDb, envData.PostgresPort)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalln("Could not initialize database", err)
	}

	return &DatabaseConnection{Database: db}
}
