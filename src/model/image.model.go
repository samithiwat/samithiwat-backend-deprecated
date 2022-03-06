package model

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ImgUrl      string         `json:"img_url"`
	OwnerID     int64          `json:"owner_id" gorm:"default:null"`
	OwnerType   string         `json:"owner_type"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
