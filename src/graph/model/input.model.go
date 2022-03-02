package model

import "time"

// TODO: Add validator

type NewSetting struct {
	ID		 int64    `json:"id"`
	AboutMe  NewAboutMe  `json:"AboutMeID"`
	Timeline NewTimeline `json:"TimelineID"`
}

type NewAboutMe struct {
	ID			int64  `json:"id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
	ImgURL      string `json:"ImgUrl"`
	SettingID   string `json:"SettingID"`
}

type NewBadge struct {
	ID		  int64   `json:"ID"`
	Name      string  `json:"Name"`
	Color     string  `json:"Color"`
	Icon      NewIcon `json:"IconID"`
	OwnerID   int     `json:"OwnerID"`
	OwnerType string  `json:"OwnerType"`
}

type NewIcon struct {
	ID		  int64  `json:"ID"`
	Name      string `json:"Name"`
	BgColor   string `json:"BgColor"`
	IconType  string `json:"IconType"`
	OwnerID   int64    `json:"OwnerID"`
	OwnerType string `json:"OwnerType"`
}

type NewImage struct {
	ID			int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgURL      string `json:"imgUrl"`
	OwnerID     string `json:"ownerId"`
	OwnerType   string `json:"ownerType"`
}

type NewTimeline struct {
	ID			int64  	   `json:"id"`
	Slug        string     `json:"Slug"`
	Name        string     `json:"Name"`
	Description string     `json:"Description"`
	Thumbnail   string     `json:"Thumbnail"`
	EventDate   time.Time  `json:"EventDate"`
	Images 		[]*NewImage`json:"Images"`
	Icon        *NewIcon   `json:"Icon"`
	SettingID   string     `json:"SettingID"`
}

type NewGithubRepo struct {
	ID			 int64  	`json:"id"`
	Name	  	 string    	`json:"name"`
	Author	  	 string    	`json:"author"`
	Description  string    	`json:"description"`
	ThumbnailUrl string    	`json:"thumbnail_url"`
	Url 	 	 string    	`json:"url"`
	LatestUpdate time.Time 	`json:"latest_update"`
	Star	 	 int64     	`json:"star"`
	Framework    NewBadge 	 	`json:"framework"`
	Language	 NewBadge 	 	`json:"language"`
}