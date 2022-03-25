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
	config     *config.Config
	connection *gorm.DB
}

func InitDatabase(config *config.Config) (Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Database.Host, strconv.Itoa(config.Database.Port), config.Database.User, config.Database.Password, config.Database.Name, config.Database.SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &database{connection: db, config: config}, nil
}

func (d *database) GetConnection() *gorm.DB {
	return d.connection
}

func (d *database) AutoMigrate() error {
	return d.connection.AutoMigrate(&model.Image{}, &model.Badge{}, &model.Icon{}, &model.Setting{}, &model.AboutMe{}, &model.Timeline{}, &model.GithubRepo{})
}
