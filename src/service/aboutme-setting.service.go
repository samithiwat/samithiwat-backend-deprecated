package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
)

type AboutMeSettingService interface {
	GetAll() ([]*model2.AboutMe, error)
	GetOne(id int64) (*model2.AboutMe, error)
	Create(settingDto *model2.NewAboutMe) (*model2.AboutMe, error)
	Update(id int64, imageDto *model2.NewAboutMe) (*model2.AboutMe, error)
	Delete(id int64) (*model2.AboutMe, error)
	DtoToRaw(settingDto *model2.NewAboutMe) (*model2.AboutMe, error)
}

func NewAboutMeSettingService(db database.Database, validatorService ValidatorService) AboutMeSettingService {
	return &aboutMeSettingService{
		database:         db,
		validatorService: validatorService,
	}
}

type aboutMeSettingService struct {
	database         database.Database
	validatorService ValidatorService
}

func (s *aboutMeSettingService) GetAll() ([]*model2.AboutMe, error) {
	db := s.database.GetConnection()

	var settings []*model2.AboutMe

	result := db.Find(&settings)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return settings, nil
}

func (s *aboutMeSettingService) GetOne(id int64) (*model2.AboutMe, error) {
	db := s.database.GetConnection()

	var setting *model2.AboutMe

	result := db.First(&setting, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return setting, nil
}

func (s *aboutMeSettingService) Create(aboutMeDto *model2.NewAboutMe) (*model2.AboutMe, error) {
	db := s.database.GetConnection()

	setting, err := s.DtoToRaw(aboutMeDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *aboutMeSettingService) Update(id int64, aboutMeDto *model2.NewAboutMe) (*model2.AboutMe, error) {
	db := s.database.GetConnection()

	var aboutMe *model2.AboutMe
	raw, err := s.DtoToRaw(aboutMeDto)
	if err != nil {
		return nil, err
	}

	result := db.First(&aboutMe, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return aboutMe, nil
}

func (s *aboutMeSettingService) Delete(id int64) (*model2.AboutMe, error) {
	db := s.database.GetConnection()

	var aboutMe *model2.AboutMe

	result := db.First(&aboutMe, id).Delete(&model2.AboutMe{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return aboutMe, nil
}

func (s *aboutMeSettingService) DtoToRaw(settingDto *model2.NewAboutMe) (*model2.AboutMe, error) {
	err := s.validatorService.AboutMe(*settingDto)
	if err != nil {
		return nil, err
	}

	aboutMe := model2.AboutMe{
		ID:          settingDto.ID,
		Name:        settingDto.Name,
		Description: settingDto.Description,
		Content:     settingDto.Content,
		ImgUrl:      settingDto.ImgURL,
		SettingID:   settingDto.SettingID,
	}

	return &aboutMe, err
}
