package model

import "time"

type NewSetting struct {
	AboutMeID  string `json:"AboutMeID"`
	TimelineID string `json:"TimelineID"`
}

type NewAboutMe struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
	ImgURL      string `json:"ImgUrl"`
	SettingID   string `json:"SettingID"`
}

type NewBadge struct {
	Name      string `json:"Name"`
	Color     string `json:"Color"`
	IconID    string `json:"IconID"`
	OwnerID   int    `json:"OwnerID"`
	OwnerType string `json:"OwnerType"`
}

type NewIcon struct {
	Name      string `json:"Name"`
	BgColor   string `json:"BgColor"`
	IconType  string `json:"IconType"`
	OwnerID   int    `json:"OwnerID"`
	OwnerType string `json:"OwnerType"`
}

type NewImage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgURL      string `json:"imgUrl"`
	OwnerID     string `json:"ownerId"`
	OwnerType   string `json:"ownerType"`
}

type NewTimeline struct {
	Slug        string    `json:"Slug"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Thumbnail   string    `json:"Thumbnail"`
	EventDate   time.Time `json:"EventDate"`
	Icon        *NewIcon  `json:"Icon"`
	SettingID   string    `json:"SettingID"`
}
