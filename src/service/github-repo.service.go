package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/database"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
)

// TODO: fetch data from github

type GithubRepoService interface {
	GetAll() ([]*model2.GithubRepo, error)
	GetOne(id int64) (*model2.GithubRepo, error)
	Create(githubRepoDto *model2.NewGithubRepo) (*model2.GithubRepo, error)
	Update(id int64, githubRepoDto *model2.NewGithubRepo) (*model2.GithubRepo, error)
	Delete(id int64) (*model2.GithubRepo, error)
	DtoToRaw(githubRepoDto model2.NewGithubRepo) (*model2.GithubRepo, error)
}

type githubRepoService struct {
	database         database.Database
	badgeService     BadgeService
	validatorService ValidatorService
}

func NewGithubRepoService(db database.Database, badgeService BadgeService, validatorService ValidatorService) GithubRepoService {
	return &githubRepoService{
		database:         db,
		badgeService:     badgeService,
		validatorService: validatorService,
	}
}

func (s githubRepoService) GetAll() ([]*model2.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo []*model2.GithubRepo

	result := db.Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").Find(&repo)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return repo, nil
}

func (s githubRepoService) GetOne(id int64) (*model2.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model2.GithubRepo

	result := db.Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").First(&repo, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return repo, nil
}

func (s githubRepoService) Create(githubRepoDto *model2.NewGithubRepo) (*model2.GithubRepo, error) {
	db := s.database.GetConnection()
	repo, err := s.DtoToRaw(*githubRepoDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&repo)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return repo, nil
}

func (s githubRepoService) Update(id int64, githubRepoDto *model2.NewGithubRepo) (*model2.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model2.GithubRepo
	raw, err := s.DtoToRaw(*githubRepoDto)
	if err != nil {
		return nil, err
	}

	result := db.Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").First(&repo, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, result.Error)
	}

	if (!cmp.Equal(raw.Framework, model2.Badge{})) {
		db.Model(&repo).Association("Framework").Replace(&raw.Framework)
		db.Model(&repo.Framework).Association("Icon").Append(&raw.Framework.Icon)
	}

	if (!cmp.Equal(raw.Language, model2.Badge{})) {
		db.Model(&repo).Association("Language").Replace(&raw.Language)
		db.Model(&repo.Language).Association("Icon").Append(&raw.Language.Icon)
	}

	return repo, nil
}

func (s githubRepoService) Delete(id int64) (*model2.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model2.GithubRepo

	result := db.First(&repo, id).Delete(&model2.GithubRepo{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return repo, nil
}

func (s githubRepoService) DtoToRaw(githubRepoDto model2.NewGithubRepo) (*model2.GithubRepo, error) {
	err := s.validatorService.GithubRepo(githubRepoDto)
	if err != nil {
		return nil, err
	}

	rawFramework, err := s.badgeService.DtoToRaw(githubRepoDto.Framework)
	if err != nil {
		return nil, err
	}

	rawLanguage, err := s.badgeService.DtoToRaw(githubRepoDto.Language)
	if err != nil {
		return nil, err
	}

	repo := model2.GithubRepo{
		ID:           githubRepoDto.ID,
		Name:         githubRepoDto.Name,
		Description:  githubRepoDto.Description,
		Author:       githubRepoDto.Author,
		ThumbnailUrl: githubRepoDto.ThumbnailUrl,
		Url:          githubRepoDto.Url,
		Star:         githubRepoDto.Star,
		LatestUpdate: githubRepoDto.LatestUpdate,
		Framework:    *rawFramework,
		Language:     *rawLanguage,
	}

	return &repo, nil
}
