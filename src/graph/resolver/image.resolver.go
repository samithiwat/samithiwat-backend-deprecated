package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
	"strconv"
	"time"

	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
)

func (r *imageResolver) CreatedDate(_ context.Context, _ *model2.Image) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) UpdatedDate(_ context.Context, _ *model2.Image) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) DeletedDate(_ context.Context, _ *model2.Image) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImage(_ context.Context, newImage model2.NewImage) (*model2.Image, error) {
	image, err := r.imageService.Create(&newImage)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (r *mutationResolver) UpdateImage(_ context.Context, id string, newImage model2.NewImage) (*model2.Image, error) {
	parsedID, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	image, err := r.imageService.Update(int64(parsedID), &newImage)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (r *mutationResolver) DeleteImage(_ context.Context, id string) (*model2.Image, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	image, err := r.imageService.Delete(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (r *queryResolver) Images(_ context.Context) ([]*model2.Image, error) {
	images, err := r.imageService.GetAll()
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (r *queryResolver) Image(_ context.Context, id string) (*model2.Image, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	image, err := r.imageService.GetOne(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return image, nil
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
