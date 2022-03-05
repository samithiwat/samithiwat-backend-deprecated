package model

import "time"

// TODO: Add validator

type NewSetting struct {
	ID          int64       `json:"id"`
	AboutMe     NewAboutMe  `json:"AboutMeID"`
	Timeline    NewTimeline `json:"TimelineID"`
	IsActivated bool        `json:"isActivated"`
}

type NewAboutMe struct {
	ID          int64  `json:"id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
	ImgURL      string `json:"ImgUrl" validate:"url"`
	SettingID   int64  `json:"SettingID"`
}

type NewBadge struct {
	ID        int64   `json:"ID"`
	Name      string  `json:"Name"`
	Color     string  `json:"Color"`
	Icon      NewIcon `json:"IconID"`
	OwnerID   int64   `json:"OwnerID"`
	OwnerType string  `json:"OwnerType"`
}

type NewIcon struct {
	ID        int64  `json:"ID"`
	Name      string `json:"Name"`
	BgColor   string `json:"BgColor" validate:"hexcolor"`
	IconType  string `json:"IconType"`
	OwnerID   int64  `json:"OwnerID"`
	OwnerType string `json:"OwnerType"`
}

type NewImage struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgURL      string `json:"imgUrl" validate:"url"`
	OwnerID     int64  `json:"ownerId"`
	OwnerType   string `json:"ownerType"`
}

type NewTimeline struct {
	ID          int64       `json:"id"`
	Slug        string      `json:"Slug"`
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Thumbnail   string      `json:"Thumbnail" validate:"url"`
	EventDate   time.Time   `json:"EventDate"`
	Images      []*NewImage `json:"Images"`
	Icon        *NewIcon    `json:"Icon"`
	SettingID   int64       `json:"SettingID"`
}

type NewGithubRepo struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Author       string    `json:"author"`
	Description  string    `json:"description"`
	ThumbnailUrl string    `json:"thumbnail_url" validate:"url"`
	Url          string    `json:"url" validate:"url"`
	LatestUpdate time.Time `json:"latest_update"`
	Star         int64     `json:"star"`
	Framework    NewBadge  `json:"framework"`
	Language     NewBadge  `json:"language"`
}
