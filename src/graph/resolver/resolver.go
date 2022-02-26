package graph

import service "github.com/samithiwat/samithiwat-backend/src/graph/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	imageService service.ImageService
}

func NewResolver(imageService service.ImageService) *Resolver {
	return &Resolver{
		imageService: imageService,
	}
}