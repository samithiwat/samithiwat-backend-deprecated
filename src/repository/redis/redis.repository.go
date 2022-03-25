package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
)

type RedisRepository struct {
	db database.Cache
}

func NewRedisRepository(db database.Cache) *RedisRepository {
	return &RedisRepository{db: db}
}
