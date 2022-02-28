//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/samithiwat/samithiwat-backend/src/database"
	graph "github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	service "github.com/samithiwat/samithiwat-backend/src/graph/services"
)

func InitializeResolver(db database.Database) (*graph.Resolver, error) {
	wire.Build(graph.NewResolver, service.NewImageService, service.NewBadgeService, service.NewIconService,service.NewSettingService , service.NewTimelineSettingService, service.NewAboutMeSettingService)
	return &graph.Resolver{}, nil
}