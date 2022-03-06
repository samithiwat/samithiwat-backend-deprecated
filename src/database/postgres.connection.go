package database

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend/src/config"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

type Database interface {
	GetConnection() *gorm.DB
	AutoMigrate() error
}

type database struct {
	connection *gorm.DB
}

func InitDatabase() (Database, error) {
	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", loadConfig.Database.Host, strconv.Itoa(loadConfig.Database.Port), loadConfig.Database.User, loadConfig.Database.Password, loadConfig.Database.Name, loadConfig.Database.SSL)

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
	return d.connection.AutoMigrate(&model.Image{}, &model.Badge{}, &model.Icon{}, &model.Setting{}, &model.AboutMe{}, &model.Timeline{}, &model.GithubRepo{})
}
