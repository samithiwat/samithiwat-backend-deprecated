package model

import (
	"time"

	"gorm.io/gorm"
)

type GithubRepo struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	Author       string         `json:"author"`
	Description  string         `json:"description"`
	ThumbnailUrl string         `json:"thumbnail_url"`
	Url          string         `json:"url"`
	LatestUpdate time.Time      `json:"latest_update"`
	Star         int64          `json:"star"`
	Framework    Badge          `json:"framework" gorm:"polymorphic:Owner;polymorphicValue:framework"`
	Language     Badge          `json:"language" gorm:"polymorphic:Owner;polymorphicValue:language"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
