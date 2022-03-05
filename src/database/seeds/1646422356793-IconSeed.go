package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"log"
)

func (s Seed) IconSeed_1646422356793() model.Icon {
	db := s.db.GetConnection()

	iconTypes := []enum.IconType{"svg", "icon"}
	idx, _ := faker.RandomInt(0, 1)

	icon := model.Icon{Name: faker.Word(), BgColor: faker.Word(), IconType: iconTypes[idx[0]]}

	result := db.Create(&icon)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	return icon
}
