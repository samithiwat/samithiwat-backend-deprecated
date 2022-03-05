package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type GithubRepoService interface {
	GetAll() ([]*model.GithubRepo, error)
	GetOne(id int64) (*model.GithubRepo, error)
	Create(githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error)
	Update(id int64, githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error)
	Delete(id int64) (*model.GithubRepo, error)
	DtoToRaw(githubRepoDto model.NewGithubRepo) *model.GithubRepo
}

type githubRepoService struct {
	database     database.Database
	badgeService BadgeService
}

func NewGithubRepoService(db database.Database, badgeService BadgeService) GithubRepoService {
	return &githubRepoService{
		database:     db,
		badgeService: badgeService,
	}
}

func (s githubRepoService) GetAll() ([]*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo []*model.GithubRepo

	result := db.Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").Find(&repo)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return repo, nil
}

func (s githubRepoService) GetOne(id int64) (*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model.GithubRepo

	result := db.Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").First(&repo, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return repo, nil
}

func (s githubRepoService) Create(githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error) {
	db := s.database.GetConnection()
	repo := s.DtoToRaw(*githubRepoDto)

	result := db.Create(&repo)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return repo, nil
}

func (s githubRepoService) Update(id int64, githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model.GithubRepo
	raw := s.DtoToRaw(*githubRepoDto)

	result := db.Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").First(&repo, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, result.Error)
	}

	if (!cmp.Equal(raw.Framework, model.Badge{})) {
		db.Model(&repo).Association("Framework").Replace(&raw.Framework)
		db.Model(&repo.Framework).Association("Icon").Append(&raw.Framework.Icon)
	}

	if (!cmp.Equal(raw.Language, model.Badge{})) {
		db.Model(&repo).Association("Language").Replace(&raw.Language)
		db.Model(&repo.Language).Association("Icon").Append(&raw.Language.Icon)
	}

	return repo, nil
}

func (s githubRepoService) Delete(id int64) (*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model.GithubRepo

	result := db.First(&repo, id).Delete(&model.GithubRepo{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return repo, nil
}

func (s githubRepoService) DtoToRaw(githubRepoDto model.NewGithubRepo) *model.GithubRepo {
	rawFramework := s.badgeService.DtoToRaw(githubRepoDto.Framework)
	rawLanguage := s.badgeService.DtoToRaw(githubRepoDto.Language)

	repo := model.GithubRepo{ID: githubRepoDto.ID, Name: githubRepoDto.Name, Description: githubRepoDto.Description, Author: githubRepoDto.Author, ThumbnailUrl: githubRepoDto.ThumbnailUrl, Url: githubRepoDto.Url, Star: githubRepoDto.Star, LatestUpdate: githubRepoDto.LatestUpdate, Framework: *rawFramework, Language: *rawLanguage}
	return &repo
}
