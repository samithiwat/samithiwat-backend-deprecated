package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
)

type TimelineSettingService interface {
	GetAll() ([]*model2.Timeline, error)
	GetOne(id int64) (*model2.Timeline, error)
	Create(settingDto *model2.NewTimeline) (*model2.Timeline, error)
	Update(id int64, imageDto *model2.NewTimeline) (*model2.Timeline, error)
	Delete(id int64) (*model2.Timeline, error)
	DtoToRaw(settingDto *model2.NewTimeline) (*model2.Timeline, error)
}

func NewTimelineSettingService(db database.Database, iconService IconService, imageService ImageService, validatorService ValidatorService) TimelineSettingService {
	return &timelineSettingService{
		database:         db,
		iconService:      iconService,
		imageService:     imageService,
		validatorService: validatorService,
	}
}

type timelineSettingService struct {
	database         database.Database
	iconService      IconService
	imageService     ImageService
	validatorService ValidatorService
}

func (s *timelineSettingService) GetAll() ([]*model2.Timeline, error) {
	db := s.database.GetConnection()

	var settings []*model2.Timeline

	result := db.Preload("Images").Preload("Icon").Find(&settings)
	if result.Error != nil {
		return nil, result.Error
	}

	return settings, nil
}

func (s *timelineSettingService) GetOne(id int64) (*model2.Timeline, error) {
	db := s.database.GetConnection()

	var setting *model2.Timeline

	result := db.Preload("Icon").Preload("Images").First(&setting, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return setting, nil
}

func (s *timelineSettingService) Create(timelineDto *model2.NewTimeline) (*model2.Timeline, error) {
	db := s.database.GetConnection()

	setting, err := s.DtoToRaw(timelineDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return setting, nil
}

func (s *timelineSettingService) Update(id int64, timelineDto *model2.NewTimeline) (*model2.Timeline, error) {
	db := s.database.GetConnection()

	var timeline *model2.Timeline
	raw, err := s.DtoToRaw(timelineDto)
	if err != nil {
		return nil, err
	}

	result := db.First(&timeline, "id = ?", id).Omit("SettingID").Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while query")
	}

	if (raw.Icon != model2.Icon{}) {
		db.Model(&timeline).Association("Icon").Replace(&raw.Icon)
	}

	if len(raw.Images) > 0 {
		db.Model(&timeline).Association("Images").Replace(raw.Images)
	}

	return timeline, nil
}

func (s *timelineSettingService) Delete(id int64) (*model2.Timeline, error) {
	db := s.database.GetConnection()

	var timeline *model2.Timeline

	result := db.First(&timeline, id).Delete(&model2.Timeline{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while query")
	}

	return timeline, nil
}

func (s *timelineSettingService) DtoToRaw(settingDto *model2.NewTimeline) (*model2.Timeline, error) {
	err := s.validatorService.Timeline(*settingDto)
	if err != nil {
		return nil, err
	}

	var icon *model2.Icon
	if settingDto.Icon != nil {
		icon, err = s.iconService.DtoToRaw(*settingDto.Icon)
		if err != nil {
			return nil, err
		}
	}

	var images []model2.Image
	for _, dto := range settingDto.Images {
		raw, err := s.imageService.DtoToRaw(*dto)
		if err != nil {
			return nil, err
		}
		images = append(images, *raw)
	}

	timeline := model2.Timeline{
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
