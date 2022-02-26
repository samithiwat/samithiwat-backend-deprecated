package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

func (r *imageResolver) CreatedDate(ctx context.Context, obj *model.Image) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) UpdatedDate(ctx context.Context, obj *model.Image) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) DeletedDate(ctx context.Context, obj *model.Image) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImage(ctx context.Context, newImage model.NewImage) (*model.Image, error) {
	image, err := r.imageService.CreateImage(&newImage)	
	if err != nil{
		return nil, err
	}

	return image, nil
}

func (r *mutationResolver) UpdateImage(ctx context.Context, id string, newImage model.NewImage) (*model.Image, error) {
	parsedID, err := strconv.Atoi(id)
	
	if err != nil {
		return nil, err
	}

	image, err := r.imageService.UpdateImage(int64(parsedID), &newImage)	
	if err != nil{
		return nil, err
	}
	
	return image, nil
}

func (r *mutationResolver) DeleteImage(ctx context.Context, id string) (*model.Image, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	
	image, err := r.imageService.DeleteImage(int64(parsedID))	
	if err != nil{
		return nil, err
	}

	return image, nil
}

func (r *queryResolver) Images(ctx context.Context) ([]*model.Image, error) {
	images, err := r.imageService.GetAllImages()
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (r *queryResolver) Image(ctx context.Context, id string) (*model.Image, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	
	image, err := r.imageService.GetImage(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return image, nil
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
