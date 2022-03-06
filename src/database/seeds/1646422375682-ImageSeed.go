package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"log"
)

func (s Seed) ImageSeed_1646422375682() model.Image {
	db := s.db.GetConnection()

	image := model.Image{Name: faker.Word(), Description: faker.Sentence(), ImgUrl: faker.URL()}

	result := db.Create(&image)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	return image
}
