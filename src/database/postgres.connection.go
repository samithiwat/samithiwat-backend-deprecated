package database

import (
	"fmt"
	"log"

	"github.com/samithiwat/samithiwat-backend/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	GetConnection() *gorm.DB
	AutoMigrate() error
}

type database struct {
	connection *gorm.DB
}

func InitDatabase() (Database, error) {
	config, err := config.LoadConfig(".")
    
    if err != nil {
        log.Fatal("cannot load config", err)
    }

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &database{connection: db}, nil
}

func (d *database) GetConnection() *gorm.DB {
	return d.connection
}

func (d *database) AutoMigrate() error {
	return d.connection.AutoMigrate()
}