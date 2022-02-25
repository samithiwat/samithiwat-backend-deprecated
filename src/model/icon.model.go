package model

import (
	"time"

	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"gorm.io/gorm"
)

type Icon struct {
	ID 		  int64 		 `json:"id"`
	Name      string 		 `json:"icon"`
	BgColor   string 		 `json:"icon_bg_color"`
	IconType  enum.IconType  `json:"icon_type"`
	OwnerID   int 			 `json:"owner_id"`
	OwnerType string 		 `json:"owner_type"`
	CreatedAt time.Time 	 `json:"created_at"`
	UpdatedAt time.Time 	 `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

}