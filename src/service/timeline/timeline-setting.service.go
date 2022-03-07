package timeline

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/icon"
	"github.com/samithiwat/samithiwat-backend/src/service/image"
)

type Repository interface {
	FindAllTimeline(*[]*model.Timeline) error
	FindOneTimeline(int64, *model.Timeline) error
	CreateTimeline(*model.Timeline) error
	UpdateTimeline(int64, *model.Timeline) error
	DeleteTimeline(int64, *model.Timeline) error
}

type Service struct {
	repository       Repository
	validatorService service.ValidatorService
	iconService      icon.Service
	imageService     image.Service
}

func NewTimelineSettingService(repository Repository, iconService icon.Service, imageService image.Service, validatorService service.ValidatorService) Service {
	return Service{
		repository:       repository,
		iconService:      iconService,
		imageService:     imageService,
		validatorService: validatorService,
	}
}

func (s *Service) GetAll() ([]*model.Timeline, error) {
	var settings []*model.Timeline

	err := s.repository.FindAllTimeline(&settings)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return settings, nil
}

func (s *Service) GetOne(id int64) (*model.Timeline, error) {
	var setting model.Timeline

	err := s.repository.FindOneTimeline(id, &setting)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *Service) Create(timelineDto *model.NewTimeline) (*model.Timeline, error) {
	setting, err := s.DtoToRaw(timelineDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateTimeline(setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return setting, nil
}

func (s *Service) Update(id int64, timelineDto *model.NewTimeline) (*model.Timeline, error) {
	timeline, err := s.DtoToRaw(timelineDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateTimeline(id, timeline)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return timeline, nil
}

func (s *Service) Delete(id int64) (*model.Timeline, error) {
	var timeline model.Timeline
	err := s.repository.DeleteTimeline(id, &timeline)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &timeline, nil
}

func (s *Service) DtoToRaw(settingDto *model.NewTimeline) (*model.Timeline, error) {
	err := s.validatorService.Timeline(*settingDto)
	if err != nil {
		return nil, err
	}

	var icon *model.Icon
	if settingDto.Icon != nil {
		icon, err = s.iconService.DtoToRaw(*settingDto.Icon)
		if err != nil {
			return nil, err
		}
	}

	var images []model.Image
	for _, dto := range settingDto.Images {
		timeline, err := s.imageService.DtoToRaw(*dto)
		if err != nil {
			return nil, err
		}
		images = append(images, *timeline)
	}

	timeline := model.Timeline{
		ID:          settingDto.ID,
		Name:        settingDto.Name,
		Description: settingDto.Description,
		Slug:        settingDto.Slug,
		Thumbnail:   settingDto.Thumbnail,
		EventDate:   settingDto.EventDate,
		Images:      images,
	}

	if icon != nil {
		timeline.Icon = *icon
	}

	return &timeline, nil
}
