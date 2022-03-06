package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockDatabase() (Database, sqlmock.Sqlmock, error) {
	var db *sql.DB
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return &database{connection: gdb}, mock, nil

}
