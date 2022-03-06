package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type TimelineSettingService interface {
	GetAll() ([]*model.Timeline, error)
	GetOne(id int64) (*model.Timeline, error)
	Create(settingDto *model.NewTimeline) (*model.Timeline, error)
	Update(id int64, imageDto *model.NewTimeline) (*model.Timeline, error)
	Delete(id int64) (*model.Timeline, error)
	DtoToRaw(settingDto *model.NewTimeline) (*model.Timeline, error)
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

func (s *timelineSettingService) GetAll() ([]*model.Timeline, error) {
	db := s.database.GetConnection()

	var settings []*model.Timeline

	result := db.Preload("Images").Preload("Icon").Find(&settings)
	if result.Error != nil {
		return nil, result.Error
	}

	return settings, nil
}

func (s *timelineSettingService) GetOne(id int64) (*model.Timeline, error) {
	db := s.database.GetConnection()

	var setting *model.Timeline

	result := db.Preload("Icon").Preload("Images").First(&setting, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return setting, nil
}

func (s *timelineSettingService) Create(timelineDto *model.NewTimeline) (*model.Timeline, error) {
	db := s.database.GetConnection()

	setting, err := s.DtoToRaw(timelineDto)
	if err != nil{
		return nil, err
	}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return setting, nil
}

func (s *timelineSettingService) Update(id int64, timelineDto *model.NewTimeline) (*model.Timeline, error) {
	db := s.database.GetConnection()

	var timeline *model.Timeline
	raw, err := s.DtoToRaw(timelineDto)
	if err != nil{
		return nil, err
	}

	result := db.First(&timeline, "id = ?", id).Omit("SettingID").Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while query")
	}

	if (raw.Icon != model.Icon{}) {
		db.Model(&timeline).Association("Icon").Replace(&raw.Icon)
	}

	if len(raw.Images) > 0 {
		db.Model(&timeline).Association("Images").Replace(raw.Images)
	}

	return timeline, nil
}

func (s *timelineSettingService) Delete(id int64) (*model.Timeline, error) {
	db := s.database.GetConnection()

	var timeline *model.Timeline

	result := db.First(&timeline, id).Delete(&model.Timeline{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while query")
	}

	return timeline, nil
}

func (s *timelineSettingService) DtoToRaw(settingDto *model.NewTimeline) (*model.Timeline, error) {
	err := s.validatorService.Timeline(*settingDto)
	if err != nil{
		return nil, err
	}

	var icon *model.Icon
	if settingDto.Icon != nil {
		icon, err = s.iconService.DtoToRaw(*settingDto.Icon)
		if err != nil{
			return nil, err
		}
	}

	var images []model.Image
	for _, dto := range settingDto.Images {
		raw, err := s.imageService.DtoToRaw(*dto)
		if err != nil{
			return nil, err
		}
		images = append(images, *raw)
	}

	timeline := model.Timeline{
		ID: settingDto.ID,
		Name: settingDto.Name,
		Description: settingDto.Description,
		Slug: settingDto.Slug,
		Thumbnail: settingDto.Thumbnail,
		EventDate: settingDto.EventDate,
		Images: images,
	}

	if icon != nil {
		timeline.Icon = *icon
	}

	return &timeline, nil
}
