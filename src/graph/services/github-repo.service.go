package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"gorm.io/gorm"
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
	database database.Database
	badgeService BadgeService
}

func NewGithubRepoService(db database.Database, badgeService BadgeService) GithubRepoService {
	return &githubRepoService{
		database: db,
		badgeService: badgeService,
	}
}

func (s githubRepoService) GetAll() ([]*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo []*model.GithubRepo

	result := db.Find(&repo)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo, nil
}

func (s githubRepoService) GetOne(id int64) (*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model.GithubRepo

	result := db.Preload("Framework").Preload("Language").First(&repo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo, nil
}

func (s githubRepoService) Create(githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error) {
	db := s.database.GetConnection()
	repo := s.DtoToRaw(*githubRepoDto)

	result := db.Create(&repo)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return repo, nil
}


func (s githubRepoService) Update(id int64, githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error) {
	// TODO: Complete this

	db := s.database.GetConnection()

	var repo *model.GithubRepo
	rawGithubRepo := s.DtoToRaw(*githubRepoDto)

	result := db.First(&repo, "id = ?", id).Updates(rawGithubRepo)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&repo)
	return repo, nil
}

func (s githubRepoService) Delete(id int64) (*model.GithubRepo, error) {
	db := s.database.GetConnection()

	var repo *model.GithubRepo

	result := db.First(&repo, id).Delete(&model.GithubRepo{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return repo, nil
}

func (s githubRepoService) DtoToRaw(githubRepoDto model.NewGithubRepo) *model.GithubRepo {
	rawFramework := s.badgeService.DtoToRaw(githubRepoDto.Framework)
	rawLanguage := s.badgeService.DtoToRaw(githubRepoDto.Language)
	repo := model.GithubRepo{Framework: *rawFramework, Language: *rawLanguage}
	return &repo
}