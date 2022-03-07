package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"log"
	"time"
)

func (s Seed) TimelineSeed1646426038041() model.Timeline {

	db := s.db.GetConnection()

	icon := s.IconSeed1646422356793()

	setting := model.Timeline{Name: faker.Word(), Description: faker.Sentence(), Slug: faker.Word(), EventDate: time.Now(), Thumbnail: faker.URL(), Icon: icon}

	result := db.Create(&setting)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	n, _ := faker.RandomInt(1, 5)

	for j := 0; j < n[0]; j++ {
		image := model.Image{Name: faker.Word(), Description: faker.Sentence(), ImgUrl: faker.URL(), OwnerID: setting.ID, OwnerType: "settings_timeline"}
		result := db.Create(&image)

		if result.Error != nil {
			log.Fatalln(result.Error)
		}
	}

	return setting
}
