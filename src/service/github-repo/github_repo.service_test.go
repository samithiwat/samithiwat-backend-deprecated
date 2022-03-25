package github_test

import (
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/badge"
	"github.com/samithiwat/samithiwat-backend/src/service/github-repo"
	"github.com/samithiwat/samithiwat-backend/src/service/icon"
	"math/rand"
	"time"
)

var latestUpdate, _ = time.Parse(faker.Timestamp(), faker.Timestamp())

var iconDto = model.NewIcon{
	Name:     faker.Word(),
	IconType: string(enum.ICON),
	BgColor:  "#fff",
}

var badgeDto = model.NewBadge{
	Name:  faker.Word(),
	Icon:  iconDto,
	Color: "#fff",
}

var dto = model.NewGithubRepo{
	Name:         faker.Word(),
	Description:  faker.Sentence(),
	Author:       faker.Username(),
	Star:         int64(rand.Intn(10000)),
	LatestUpdate: latestUpdate,
	Url:          faker.URL(),
	ThumbnailUrl: faker.URL(),
	Language:     badgeDto,
	Framework:    badgeDto,
}

var repo1 = model.GithubRepo{}
var repo2 = model.GithubRepo{}
var repo3 = model.GithubRepo{}
var repo4 = model.GithubRepo{}
var repo5 = model.GithubRepo{}
var repos []*model.GithubRepo

type DummyIconDB struct{}

func (DummyIconDB) FindAllIcon(*[]*model.Icon) error {
	return nil
}
func (DummyIconDB) FindOneIcon(int64, *model.Icon) error {
	return nil
}
func (DummyIconDB) CreateIcon(*model.Icon) error {
	return nil
}
func (DummyIconDB) UpdateIcon(int64, *model.Icon) error {
	return nil
}
func (DummyIconDB) DeleteIcon(int64, *model.Icon) error {
	return nil
}

type DummyBadgeDB struct{}

func (DummyBadgeDB) FindAllBadge(*[]*model.Badge) error {
	return nil
}
func (DummyBadgeDB) FindOneBadge(int64, *model.Badge) error {
	return nil
}
func (DummyBadgeDB) CreateBadge(*model.Badge) error {
	return nil
}
func (DummyBadgeDB) UpdateBadge(int64, *model.Badge) error {
	return nil
}
func (DummyBadgeDB) DeleteBadge(int64, *model.Badge) error {
	return nil
}

type TestDB struct{}

func (TestDB) FindAllGithubRepo(ghRepos *[]*model.GithubRepo) error {
	*ghRepos = repos
	return nil
}

func (TestDB) FindOneGithubRepo(id int64, repo *model.GithubRepo) error {
	*repo = *repos[id-1]
	return nil
}

func (TestDB) CreateGithubRepo(repo *model.GithubRepo) error {
	*repo = repo5
	return nil
}

func (TestDB) UpdateGithubRepo(_ int64, repo *model.GithubRepo) error {
	*repo = repo5
	return nil
}

func (TestDB) DeleteGithubRepo(_ int64, repo *model.GithubRepo) error {
	*repo = repo1
	return nil
}

type TestErrorDB struct{}

func (TestErrorDB) FindAllGithubRepo(repo *[]*model.GithubRepo) error {
	var result []*model.GithubRepo
	*repo = result
	return nil
}

func (TestErrorDB) FindOneGithubRepo(_ int64, repo *model.GithubRepo) error {
	repo = &model.GithubRepo{}
	return nil
}

func (TestErrorDB) CreateGithubRepo(*model.GithubRepo) error {
	return errors.New("duplicated")
}

func (TestErrorDB) UpdateGithubRepo(int64, *model.GithubRepo) error {
	return errors.New("not found")
}

func (TestErrorDB) DeleteGithubRepo(int64, *model.GithubRepo) error {
	return errors.New("not found")
}

type TestCache struct{}

func (TestCache) GetGithubRepoDetails() error {
	return nil
}

var _ = Describe("GithubRepo.Service", func() {
	var repoService github.Service
	var validator service.ValidatorService
	var iconService icon.Service
	var badgeService badge.Service

	BeforeEach(func() {
		validator = service.NewValidatorService()
		iconService = icon.NewIconService(DummyIconDB{}, validator)
		badgeService = badge.NewBadgeService(DummyBadgeDB{}, iconService, validator)
		err := faker.FakeData(&repo1)
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = faker.FakeData(&repo2)
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = faker.FakeData(&repo3)
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = faker.FakeData(&repo4)
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = faker.FakeData(&repo5)
		if err != nil {
			fmt.Printf("%v", err)
		}

		repos = append(repos, &repo1, &repo2, &repo3, &repo4)
	})

	Context("FindAll", func() {
		Describe("Not Empty", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestDB{}, TestCache{}, badgeService, validator)
			})

			It("Should return all repo", func() {
				want := repos
				result, err := repoService.GetAll()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(want))
			})
		})
		Describe("Empty", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestErrorDB{}, TestCache{}, badgeService, validator)
			})

			It("Should return all repo", func() {
				var want []*model.GithubRepo
				result, err := repoService.GetAll()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(want))
			})
		})
	})

	Context("FindOne", func() {
		Describe("Not Empty", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestDB{}, TestCache{}, badgeService, validator)
			})

			It("Should return repo1", func() {
				want := repo1
				result, err := repoService.GetOne(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(want))
			})

			It("Should return repo2", func() {
				want := repo2
				result, err := repoService.GetOne(2)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(want))
			})

			It("Should return repo3", func() {
				want := repo3
				result, err := repoService.GetOne(3)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(want))
			})

			It("Should return repo4", func() {
				want := repo4
				result, err := repoService.GetOne(4)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(want))
			})
		})
		Describe("Empty", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestErrorDB{}, TestCache{}, badgeService, validator)
			})

			It("Should return empty model", func() {
				want := model.GithubRepo{}
				result, err := repoService.GetOne(100)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(want))
			})
		})
	})

	Context("Create", func() {
		Describe("Not Error", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestDB{}, TestCache{}, badgeService, validator)
			})

			It("Should create new repo", func() {
				want := repo5
				result, err := repoService.Create(&dto)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(want))
			})
		})

		Describe("Error", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestErrorDB{}, TestCache{}, badgeService, validator)
			})

			It("Should throw Bad Request error", func() {
				dto := model.NewGithubRepo{}

				err := faker.FakeData(&dto)
				if err != nil {
					fmt.Printf("%v", err)
				}

				result, err := repoService.Create(&dto)

				Expect(err).Should(HaveOccurred())
				Expect(err.(*fiber.Error).Code).Should(Equal(fiber.StatusBadRequest))
				Expect(result).Should(BeNil())
			})

			It("Should throw Unprocessable Entity error", func() {
				result, err := repoService.Create(&dto)

				Expect(err).Should(HaveOccurred())
				Expect(err.(*fiber.Error).Code).Should(Equal(fiber.StatusUnprocessableEntity))
				Expect(result).Should(BeNil())
			})
		})
	})

	Context("Update", func() {
		Describe("Not Error", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestDB{}, TestCache{}, badgeService, validator)
			})

			It("Should call update repo", func() {
				result, err := repoService.Update(1, &dto)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(repo5))
			})
		})

		Describe("Error", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestErrorDB{}, TestCache{}, badgeService, validator)
			})

			It("Should throw Bad Request error", func() {
				dto := model.NewGithubRepo{}

				err := faker.FakeData(&dto)
				if err != nil {
					fmt.Printf("%v", err)
				}

				result, err := repoService.Update(1, &dto)

				Expect(err).Should(HaveOccurred())
				Expect(err.(*fiber.Error).Code).Should(Equal(fiber.StatusBadRequest))
				Expect(result).Should(BeNil())
			})

			It("Should throw Not Found error", func() {
				result, err := repoService.Update(100, &dto)

				Expect(err).Should(HaveOccurred())
				Expect(err.(*fiber.Error).Code).Should(Equal(fiber.StatusNotFound))
				Expect(result).Should(BeNil())
			})
		})
	})

	Context("Delete", func() {
		Describe("Not Error", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestDB{}, TestCache{}, badgeService, validator)
			})

			It("Should call delete repo", func() {
				result, err := repoService.Delete(1)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(*result).Should(Equal(repo1))
			})
		})

		Describe("Error", func() {
			BeforeEach(func() {
				repoService = github.NewGithubRepoService(TestErrorDB{}, TestCache{}, badgeService, validator)
			})

			It("Should throw Not Found error", func() {
				result, err := repoService.Delete(100)

				Expect(err).Should(HaveOccurred())
				Expect(err.(*fiber.Error).Code).Should(Equal(fiber.StatusNotFound))
				Expect(result).Should(BeNil())
			})
		})
	})
})
