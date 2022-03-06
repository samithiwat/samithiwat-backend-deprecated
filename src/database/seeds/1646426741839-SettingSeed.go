package seed

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (s Seed) SettingSeed1646426741839() model.Setting {

	db := s.db.GetConnection()

	timeline := s.TimelineSeed1646426038041()
	aboutMe := s.AboutMeSeed1646425792323()

	setting := model.Setting{AboutMe: aboutMe, Timeline: timeline}

	db.Create(&setting)

	return setting
}
