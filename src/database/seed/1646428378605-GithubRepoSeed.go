package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"time"
)

func (s Seed) GithubRepoSeed_1646428378605() model.GithubRepo {

	db := s.db.GetConnection()

	language := s.BadgeSeed_1646422394617()
	framework := s.BadgeSeed_1646422394617()

	stars, _ := faker.RandomInt(0, 1000000)

	repo := model.GithubRepo{Name: faker.Word(), Description: faker.Sentence(), Author: faker.Username(), Star: int64(stars[0]), ThumbnailUrl: faker.URL(), Url: faker.URL(), Language: language, Framework: framework, LatestUpdate: time.Now()}

	db.Create(&repo)

	return repo
}
