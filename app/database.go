package app

import (
	"fmt"
	"os"
	"task/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func NewDB(config *ConfigDB) *gorm.DB {
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicHandling(err)
	fmt.Printf("database Successfully connected!")
	return db
}
