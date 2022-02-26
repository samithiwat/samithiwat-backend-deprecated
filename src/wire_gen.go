// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	"github.com/samithiwat/samithiwat-backend/src/graph/services"
)

// Injectors from wire.go:

func InitializeResolver(db database.Database) (*graph.Resolver, error) {
	imageService := service.NewImageService(db)
	resolver := graph.NewResolver(imageService)
	return resolver, nil
}
