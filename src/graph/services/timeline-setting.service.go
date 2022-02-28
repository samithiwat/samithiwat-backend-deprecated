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
}

func NewTimelineSettingService(db database.Database) TimelineSettingService {
	return &timelineSettingService{
		database: db,
	}
}

type timelineSettingService struct {
	database database.Database
}

func (s *timelineSettingService) GetAll() ([]*model.Timeline, error){
	db := s.database.GetConnection()

	var settings []*model.Timeline

	result := db.Find(&settings)
	if result.Error != nil {
		return nil, result.Error
	}

	return settings, nil
}

func (s *timelineSettingService) GetOne(id int64) (*model.Timeline, error){
	db := s.database.GetConnection()

	var setting *model.Timeline

	result := db.Preload("Icon").Preload("Images").First(&setting, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return setting, nil
}

func (s *timelineSettingService) Create(timelineDto *model.NewTimeline) (*model.Timeline, error){
	db := s.database.GetConnection()

	setting := model.Timeline{Name: timelineDto.Name, Description: timelineDto.Description,Slug: timelineDto.Slug,Thumbnail: timelineDto.Thumbnail, EventDate: timelineDto.EventDate}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &setting, nil
}

func (s *timelineSettingService) Update(id int64, timelineDto *model.NewTimeline) (*model.Timeline, error) {
	db := s.database.GetConnection()

	var timeline *model.Timeline

	result := db.First(&timeline, "id = ?", id).Updates(model.Timeline{Name: timelineDto.Name, Description: timelineDto.Description,Slug: timelineDto.Slug,Thumbnail: timelineDto.Thumbnail, EventDate: timelineDto.EventDate})

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

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