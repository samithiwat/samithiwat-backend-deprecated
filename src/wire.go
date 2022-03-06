//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/samithiwat/samithiwat-backend/src/database"
	graph "github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	"github.com/samithiwat/samithiwat-backend/src/service"
)

func InitializeResolver(db database.Database) (*graph.Resolver, error) {
	wire.Build(graph.NewResolver, service.NewImageService, service.NewGithubRepoService, service.NewBadgeService, service.NewIconService, service.NewSettingService, service.NewTimelineSettingService, service.NewAboutMeSettingService, service.NewValidatorService, repository.NewGormRepository)
	return &graph.Resolver{}, nil
}
