// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	gorm "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	redis "github.com/samithiwat/samithiwat-backend/src/repository/redis"
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

func InitializeResolver(db database.Database, cache database.Cache) (*graph.Resolver, error) {
	gormRepository := gorm.NewGormRepository(db)
	redisRepository := redis.NewRedisRepository(cache)
	validatorService := service.NewValidatorService()
	imageService := image.NewImageService(gormRepository, validatorService)
	iconService := icon.NewIconService(gormRepository, validatorService)
	badgeService := badge.NewBadgeService(gormRepository, iconService, validatorService)
	aboutmeService := aboutme.NewAboutMeSettingService(gormRepository, validatorService)
	timelineService := timeline.NewTimelineSettingService(gormRepository, iconService, imageService, validatorService)
	settingService := setting.NewSettingService(gormRepository, aboutmeService, timelineService, validatorService)
	githubService := github.NewGithubRepoService(gormRepository, redisRepository, badgeService, validatorService)
	resolver := graph.NewResolver(imageService, iconService, badgeService, aboutmeService, timelineService, settingService, githubService)
	return resolver, nil
}
