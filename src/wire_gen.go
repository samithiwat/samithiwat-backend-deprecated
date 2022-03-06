// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	"github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/aboutme"
	"github.com/samithiwat/samithiwat-backend/src/service/badge"
	"github.com/samithiwat/samithiwat-backend/src/service/github-repo"
	"github.com/samithiwat/samithiwat-backend/src/service/icon"
	"github.com/samithiwat/samithiwat-backend/src/service/image"
	"github.com/samithiwat/samithiwat-backend/src/service/setting"
	"github.com/samithiwat/samithiwat-backend/src/service/timeline"
)

// Injectors from wire.go:

func InitializeResolver(db database.Database) (*graph.Resolver, error) {
	gormRepository := repository.NewGormRepository(db)
	validatorService := service.NewValidatorService()
	imageService := image.NewImageService(gormRepository, validatorService)
	iconService := icon.NewIconService(gormRepository, validatorService)
	badgeService := badge.NewBadgeService(gormRepository, iconService, validatorService)
	aboutMeSettingService := aboutme.NewAboutMeSettingService(gormRepository, validatorService)
	timelineSettingService := timeline.NewTimelineSettingService(gormRepository, iconService, imageService, validatorService)
	settingService := setting.NewSettingService(gormRepository, aboutMeSettingService, timelineSettingService, validatorService)
	githubRepoService := github.NewGithubRepoService(gormRepository, badgeService, validatorService)
	resolver := graph.NewResolver(imageService, iconService, badgeService, aboutMeSettingService, timelineSettingService, settingService, githubRepoService)
	return resolver, nil
}
