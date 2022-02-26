package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
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

func (r *mutationResolver) CreateImage(ctx context.Context, name string, description *string, imgURL string, ownerID *string, ownerType *string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateImage(ctx context.Context, id string, name *string, description *string, imgURL *string, ownerID *string, ownerType *string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteImage(ctx context.Context, id string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Images(ctx context.Context) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Image(ctx context.Context, id string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
