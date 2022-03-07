package model

type Tabler interface {
	TableName() string
}

func (AboutMe) TableName() string {
	return "settings_about_me"
}

func (Timeline) TableName() string {
	return "settings_timeline"
}
