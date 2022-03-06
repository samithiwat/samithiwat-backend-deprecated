package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

type ValidatorService interface {
	Setting(settingDto model.NewSetting) error
	Timeline(timelineDto model.NewTimeline) error
	AboutMe(aboutMeDto model.NewAboutMe) error
	GithubRepo(githubRepoDto model.NewGithubRepo) error
	Image(imageDto model.NewImage) error
	Badge(badgeDto model.NewBadge) error
	Icon(iconDto model.NewIcon) error
}

func NewValidatorService() ValidatorService {
	return &validatorService{}
}

type validatorService struct{}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func checkError(error error) error {
	var errors []string
	if error != nil {
		errors = append(errors, "Bad Request")
		for _, err := range error.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			message := fmt.Sprintf("%v", element)
			errors = append(errors, message)
		}
	}

	if errors != nil {
		return fiber.NewError(fiber.StatusBadRequest, errors)
	}

	return nil
}

func (s *validatorService) Setting(settingDto model.NewSetting) error {
	err := validate.Struct(settingDto)
	return checkError(err)
}

func (s *validatorService) AboutMe(aboutMeDto model.NewAboutMe) error {
	err := validate.Struct(aboutMeDto)
	return checkError(err)
}

func (s *validatorService) Timeline(timelineDto model.NewTimeline) error {
	err := validate.Struct(timelineDto)
	return checkError(err)
}

func (s *validatorService) GithubRepo(githubRepoDto model.NewGithubRepo) error {
	err := validate.Struct(githubRepoDto)
	return checkError(err)
}

func (s *validatorService) Image(imageDto model.NewImage) error {
	err := validate.Struct(imageDto)
	return checkError(err)
}

func (s *validatorService) Badge(badgeDto model.NewBadge) error {
	err := validate.Struct(badgeDto)
	return checkError(err)
}

func (s *validatorService) Icon(iconDto model.NewIcon) error {
	err := validate.Struct(iconDto)
	return checkError(err)
}
