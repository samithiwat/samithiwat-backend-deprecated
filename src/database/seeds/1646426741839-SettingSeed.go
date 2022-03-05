package seed

import (
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

func (s Seed) SettingSeed_1646426741839() model.Setting {

	db := s.db.GetConnection()

	timeline := s.TimelineSeed_1646426038041()
	aboutMe := s.AboutMeSeed_1646425792323()

	setting := model.Setting{AboutMe: aboutMe, Timeline: timeline}

	db.Create(&setting)

	return setting
}