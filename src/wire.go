//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/samithiwat/samithiwat-backend/src/database"
	graph "github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/aboutme"
	"github.com/samithiwat/samithiwat-backend/src/service/badge"
	"github.com/samithiwat/samithiwat-backend/src/service/github-repo"
	"github.com/samithiwat/samithiwat-backend/src/service/icon"
	"github.com/samithiwat/samithiwat-backend/src/service/image"
	"github.com/samithiwat/samithiwat-backend/src/service/setting"
	"github.com/samithiwat/samithiwat-backend/src/service/timeline"
)

func InitializeResolver(db database.Database) (*graph.Resolver, error) {
	wire.Build(graph.NewResolver, image.NewImageService, github.NewGithubRepoService, badge.NewBadgeService, icon.NewIconService, setting.NewSettingService, timeline.NewTimelineSettingService, aboutme.NewAboutMeSettingService, service.NewValidatorService, repository.NewGormRepository)
	return &graph.Resolver{}, nil
}
