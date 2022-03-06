package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
)

type GormRepository struct {
	db database.Database
}

func NewGormRepository(db database.Database) GormRepository {
	return GormRepository{db: db}
}
