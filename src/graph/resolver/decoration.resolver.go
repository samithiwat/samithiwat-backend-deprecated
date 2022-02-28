package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
	iconType, err := r.iconService.CheckIconType(obj.IconType)
	if err != nil{
		return "", err
	}

	return iconType, nil
}

func (r *iconResolver) DeletedAt(ctx context.Context, obj *model.Icon) (*time.Time, error) {
	if !obj.DeletedAt.Time.IsZero() {
		return nil, fiber.ErrNotFound
	}

	return &obj.DeletedAt.Time, nil
}

func (r *mutationResolver) CreateIcon(ctx context.Context, newIcon model.NewIcon) (*model.Icon, error) {
	icon, err := r.iconService.Create(newIcon)
	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *mutationResolver) UpdateIcon(ctx context.Context, id string, newIcon model.NewIcon) (*model.Icon, error) {
	convertedID, err := strconv.Atoi(id)

	icon, err := r.iconService.Update(int64(convertedID), newIcon)
	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *mutationResolver) DeleteIcon(ctx context.Context, id string) (*model.Icon, error) {
	convertedID, err := strconv.Atoi(id)

	icon, err := r.iconService.Delete(int64(convertedID))

	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *mutationResolver) CreateBadge(ctx context.Context, newBadge *model.NewBadge) (*model.Badge, error) {
	badge, err := r.badgeService.Create(newBadge)
	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *mutationResolver) UpdateBadge(ctx context.Context, id string, newBadge *model.NewBadge) (*model.Badge, error) {
	convertedID, err := strconv.Atoi(id)

	badge, err := r.badgeService.Update(int64(convertedID), newBadge)
	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *mutationResolver) DeleteBadge(ctx context.Context, id string) (*model.Badge, error) {
	convertedID, err := strconv.Atoi(id)

	badge, err := r.badgeService.Delete(int64(convertedID))

	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *queryResolver) Icons(ctx context.Context) ([]*model.Icon, error) {
	icons, err := r.iconService.GetAll()

	if err != nil {
		return nil, err
	}

	return icons, nil
}

func (r *queryResolver) Icon(ctx context.Context, id string) (*model.Icon, error) {
	convertedID, err := strconv.Atoi(id)

	icon, err := r.iconService.GetOne(int64(convertedID))

	if err != nil {
		return nil, err
	}

	return icon, nil
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
	badges, err := r.badgeService.GetAll()

	if err != nil {
		return nil, err
	}

	return badges, nil
}

func (r *queryResolver) Badge(ctx context.Context, id string) (*model.Badge, error) {
	convertedID, err := strconv.Atoi(id)

	badge, err := r.badgeService.GetOne(int64(convertedID))

	if err != nil {
		return nil, err
	}

	return badge, nil
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
