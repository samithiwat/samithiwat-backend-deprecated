package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/samithiwat/samithiwat-backend/src/service/aboutme"
	"github.com/samithiwat/samithiwat-backend/src/service/badge"
	"github.com/samithiwat/samithiwat-backend/src/service/github-repo"
	"github.com/samithiwat/samithiwat-backend/src/service/icon"
	"github.com/samithiwat/samithiwat-backend/src/service/image"
	"github.com/samithiwat/samithiwat-backend/src/service/setting"
	"github.com/samithiwat/samithiwat-backend/src/service/timeline"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	imageService           image.Service
	iconService            icon.Service
	badgeService           badge.Service
	aboutMeSettingService  aboutme.Service
	timelineSettingService timeline.Service
	settingService         setting.Service
	githubRepoService      github.Service
}

func NewResolver(imageService image.Service, iconService icon.Service, badgeService badge.Service, aboutMeSettingService aboutme.Service, timelineSettingService timeline.Service, settingService setting.Service, githubRepoService github.Service) *Resolver {
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
