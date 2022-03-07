package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Setting struct {
	ID          int64          `json:"id"`
	IsActivated bool           `json:"is_activated" gorm:"default:false"`
	AboutMe     AboutMe        `json:"about_me" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Timeline    Timeline       `json:"timeline" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type AboutMe struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Content     string         `json:"content"`
	ImgUrl      string         `json:"img_url"`
	SettingID   int64          `json:"setting_id" gorm:"default:null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type Timeline struct {
	ID          int64          `json:"id"`
	Slug        string         `json:"slug"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Thumbnail   string         `json:"thumbnail"`
	Images      []Image        `json:"images" gorm:"polymorphic:Owner;"`
	EventDate   time.Time      `json:"event_date"`
	Icon        Icon           `json:"icon" gorm:"polymorphic:Owner;"`
	SettingID   sql.NullInt64  `json:"setting_id" gorm:"default:null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
