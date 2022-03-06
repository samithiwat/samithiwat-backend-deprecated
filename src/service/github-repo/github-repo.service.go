package github

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/badge"
)

// TODO: fetch data from github

type Service interface {
	GetAll() ([]*model.GithubRepo, error)
	GetOne(id int64) (*model.GithubRepo, error)
	Create(githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error)
	Update(id int64, githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error)
	Delete(id int64) (*model.GithubRepo, error)
	DtoToRaw(githubRepoDto model.NewGithubRepo) (*model.GithubRepo, error)
}

type githubRepoService struct {
	repository       repository.GormRepository
	badgeService     badge.Service
	validatorService service.ValidatorService
}

func NewGithubRepoService(repository repository.GormRepository, badgeService badge.Service, validatorService service.ValidatorService) Service {
	return &githubRepoService{
		repository:       repository,
		badgeService:     badgeService,
		validatorService: validatorService,
	}
}

func (s githubRepoService) GetAll() ([]*model.GithubRepo, error) {
	var repos []*model.GithubRepo

	err := s.repository.FindAllGithubRepo(&repos)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return repos, nil
}

func (s githubRepoService) GetOne(id int64) (*model.GithubRepo, error) {
	var repo model.GithubRepo

	err := s.repository.FindGithubRepo(id, &repo)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &repo, nil
}

func (s githubRepoService) Create(githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error) {
	repo, err := s.DtoToRaw(*githubRepoDto)

	err = s.repository.CreateGithubRepo(repo)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return repo, nil
}

func (s githubRepoService) Update(id int64, githubRepoDto *model.NewGithubRepo) (*model.GithubRepo, error) {
	repo, err := s.DtoToRaw(*githubRepoDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateGithubRepo(id, repo)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return repo, nil
}

func (s githubRepoService) Delete(id int64) (*model.GithubRepo, error) {
	var repo model.GithubRepo
	err := s.repository.DeleteGithubRepo(id, &repo)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &repo, nil
}

func (s githubRepoService) DtoToRaw(githubRepoDto model.NewGithubRepo) (*model.GithubRepo, error) {
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

	repo := model.GithubRepo{
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
