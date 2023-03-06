package config

import (
	"fmt"

	"github.com/go-jwt-test/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *Config) *gorm.DB {

	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("connected successfully to the database")

	return db

}
