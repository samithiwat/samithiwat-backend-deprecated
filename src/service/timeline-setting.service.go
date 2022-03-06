package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
)

type TimelineSettingService interface {
	GetAll() ([]*model.Timeline, error)
	GetOne(id int64) (*model.Timeline, error)
	Create(settingDto *model.NewTimeline) (*model.Timeline, error)
	Update(id int64, imageDto *model.NewTimeline) (*model.Timeline, error)
	Delete(id int64) (*model.Timeline, error)
	DtoToRaw(settingDto *model.NewTimeline) (*model.Timeline, error)
}

func NewTimelineSettingService(repository repository.GormRepository, iconService IconService, imageService ImageService, validatorService ValidatorService) TimelineSettingService {
	return &timelineSettingService{
		repository:       repository,
		iconService:      iconService,
		imageService:     imageService,
		validatorService: validatorService,
	}
}

type timelineSettingService struct {
	repository       repository.GormRepository
	iconService      IconService
	imageService     ImageService
	validatorService ValidatorService
}

func (s *timelineSettingService) GetAll() ([]*model.Timeline, error) {
	var settings []*model.Timeline

	err := s.repository.FindAllTimeline(&settings)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return settings, nil
}

func (s *timelineSettingService) GetOne(id int64) (*model.Timeline, error) {
	var setting model.Timeline

	err := s.repository.FindTimeline(id, &setting)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *timelineSettingService) Create(timelineDto *model.NewTimeline) (*model.Timeline, error) {
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

func (s *timelineSettingService) Update(id int64, timelineDto *model.NewTimeline) (*model.Timeline, error) {
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

func (s *timelineSettingService) Delete(id int64) (*model.Timeline, error) {
	var timeline model.Timeline
	err := s.repository.DeleteTimeline(id, &timeline)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &timeline, nil
}

func (s *timelineSettingService) DtoToRaw(settingDto *model.NewTimeline) (*model.Timeline, error) {
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
