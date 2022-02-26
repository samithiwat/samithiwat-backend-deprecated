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

func (r *badgeResolver) Icon(ctx context.Context, obj *model.Badge) (*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *badgeResolver) DeletedAt(ctx context.Context, obj *model.Badge) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *iconResolver) IconType(ctx context.Context, obj *model.Icon) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *iconResolver) DeletedAt(ctx context.Context, obj *model.Icon) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateIcon(ctx context.Context, name string, bgColor string, iconType string, ownerID int, ownerType string) (*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateIcon(ctx context.Context, id string, name string, bgColor string, iconType string, ownerID int, ownerType string) (*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteIcon(ctx context.Context, id string) (*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBadge(ctx context.Context, name string, color string, iconID string, ownerID int, ownerType string) (*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBadge(ctx context.Context, id string, name string, color string, iconID string, ownerID int, ownerType string) (*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBadge(ctx context.Context, id string) (*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Icons(ctx context.Context) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Icon(ctx context.Context, id string) (*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IconsByOwner(ctx context.Context, ownerID int, ownerType string) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IconsByOwnerAndType(ctx context.Context, ownerID int, ownerType string, iconType string) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IconsByType(ctx context.Context, iconType string) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Badges(ctx context.Context) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Badge(ctx context.Context, id string) (*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BadgesByOwner(ctx context.Context, ownerID int, ownerType string) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BadgesByOwnerAndType(ctx context.Context, ownerID int, ownerType string, iconType string) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BadgesByType(ctx context.Context, iconType string) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

// Badge returns generated.BadgeResolver implementation.
func (r *Resolver) Badge() generated.BadgeResolver { return &badgeResolver{r} }

// Icon returns generated.IconResolver implementation.
func (r *Resolver) Icon() generated.IconResolver { return &iconResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type badgeResolver struct{ *Resolver }
type iconResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
