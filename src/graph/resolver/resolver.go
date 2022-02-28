package graph

import service "github.com/samithiwat/samithiwat-backend/src/graph/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	imageService service.ImageService
	iconService service.IconService
	badgeService service.BadgeService
	aboutMeSettingService service.AboutMeSettingService
	timelineSettingService service.TimelineSettingService
	settingService service.SettingService
}

func NewResolver(imageService service.ImageService, iconService service.IconService, badgeService service.BadgeService, aboutMeSettingService service.AboutMeSettingService, timelineSettingService service.TimelineSettingService, settingService service.SettingService) *Resolver {
	return &Resolver{
		imageService: imageService,
		iconService: iconService,
		badgeService: badgeService,
		aboutMeSettingService: aboutMeSettingService,
		timelineSettingService: timelineSettingService,
		settingService: settingService,
	}
}