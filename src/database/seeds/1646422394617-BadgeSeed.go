package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"log"
)

func (s Seed) BadgeSeed_1646422394617() model.Badge {
	db := s.db.GetConnection()

	icon := s.IconSeed_1646422356793()

	badge := model.Badge{Name: faker.Word(), Color: faker.Word(), Icon: icon}

	result := db.Create(&badge)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	return badge
}
