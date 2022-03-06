package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/samithiwat/samithiwat-backend/src/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	imageService           service.ImageService
	iconService            service.IconService
	badgeService           service.BadgeService
	aboutMeSettingService  service.AboutMeSettingService
	timelineSettingService service.TimelineSettingService
	settingService         service.SettingService
	githubRepoService      service.GithubRepoService
}

func NewResolver(imageService service.ImageService, iconService service.IconService, badgeService service.BadgeService, aboutMeSettingService service.AboutMeSettingService, timelineSettingService service.TimelineSettingService, settingService service.SettingService, githubRepoService service.GithubRepoService) *Resolver {
	return &Resolver{
		imageService:           imageService,
		iconService:            iconService,
		badgeService:           badgeService,
		aboutMeSettingService:  aboutMeSettingService,
		timelineSettingService: timelineSettingService,
		settingService:         settingService,
		githubRepoService:      githubRepoService,
	}
}
