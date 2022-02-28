package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"gorm.io/gorm"
)

type TimelineSettingService interface {
	GetAll() ([]*model.Timeline, error)
	GetOne(id int64) (*model.Timeline, error)
	Create(settingDto *model.NewTimeline) (*model.Timeline, error)
	Update(id int64, imageDto *model.NewTimeline) (*model.Timeline, error)
	Delete(id int64) (*model.Timeline, error)
	DtoToRaw(settingDto *model.NewTimeline) *model.Timeline
}

func NewTimelineSettingService(db database.Database, iconService IconService, imageService ImageService) TimelineSettingService {
	return &timelineSettingService{
		database:     db,
		iconService:  iconService,
		imageService: imageService,
	}
}

type timelineSettingService struct {
	database     database.Database
	iconService  IconService
	imageService ImageService
}

func (s *timelineSettingService) GetAll() ([]*model.Timeline, error) {
	db := s.database.GetConnection()

	var settings []*model.Timeline

	result := db.Find(&settings)
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

	setting := s.DtoToRaw(timelineDto)

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return setting, nil
}

func (s *timelineSettingService) Update(id int64, timelineDto *model.NewTimeline) (*model.Timeline, error) {
	// TODO: Make model to update with association

	db := s.database.GetConnection()

	var timeline *model.Timeline
	raw := s.DtoToRaw(timelineDto)

	result := db.Omit("SettingID").First(&timeline, "id = ?", id).Updates(raw)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	if timelineDto.Icon.ID != 0 {
		icon, err := s.iconService.GetOne(timelineDto.Icon.ID)
		if err != nil {
			return nil, err
		}

		timeline.Icon = *icon
	}

	timeline.Images = raw.Images

	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&timeline)

	return timeline, nil
}

func (s *timelineSettingService) Delete(id int64) (*model.Timeline, error) {
	db := s.database.GetConnection()

	var timeline *model.Timeline

	result := db.First(&timeline, id).Delete(&model.Timeline{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return timeline, nil
}

func (s *timelineSettingService) DtoToRaw(settingDto *model.NewTimeline) *model.Timeline {
	icon := s.iconService.DtoToRaw(*settingDto.Icon)

	var images []model.Image
	for _, dto := range settingDto.Images {
		raw := s.imageService.DtoToRaw(*dto)
		images = append(images, *raw)
	}

	timeline := model.Timeline{Icon: *icon, Images: images, ID: settingDto.ID, Name: settingDto.Name, Description: settingDto.Description, Slug: settingDto.Slug, Thumbnail: settingDto.Thumbnail, EventDate: settingDto.EventDate}

	return &timeline
}
