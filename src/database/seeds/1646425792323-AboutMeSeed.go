package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"log"
)

func (s Seed) AboutMeSeed_1646425792323() model.AboutMe {
	db := s.db.GetConnection()

	setting := model.AboutMe{Name: faker.Word(), Description: faker.Sentence(), ImgUrl: faker.URL(), Content: faker.Paragraph()}

	result := db.Create(&setting)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	return setting
}
