package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

func (r *badgeResolver) DeletedAt(_ context.Context, _ *model.Badge) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *iconResolver) IconType(_ context.Context, obj *model.Icon) (string, error) {
	iconType, err := r.iconService.CheckIconType(obj.IconType)
	if err != nil {
		return "", err
	}

	return iconType, nil
}

func (r *iconResolver) DeletedAt(_ context.Context, obj *model.Icon) (*time.Time, error) {
	if !obj.DeletedAt.Time.IsZero() {
		return nil, fiber.ErrNotFound
	}

	return &obj.DeletedAt.Time, nil
}

func (r *mutationResolver) CreateIcon(_ context.Context, newIcon model.NewIcon) (*model.Icon, error) {
	icon, err := r.iconService.Create(newIcon)
	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *mutationResolver) UpdateIcon(_ context.Context, id string, newIcon model.NewIcon) (*model.Icon, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err)
	}

	icon, err := r.iconService.Update(int64(parsedID), newIcon)
	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *mutationResolver) DeleteIcon(_ context.Context, id string) (*model.Icon, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err)
	}

	icon, err := r.iconService.Delete(int64(parsedID))

	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *mutationResolver) CreateBadge(_ context.Context, newBadge *model.NewBadge) (*model.Badge, error) {
	badge, err := r.badgeService.Create(newBadge)
	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *mutationResolver) UpdateBadge(_ context.Context, id string, newBadge *model.NewBadge) (*model.Badge, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err)
	}

	badge, err := r.badgeService.Update(int64(parsedID), newBadge)
	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *mutationResolver) DeleteBadge(_ context.Context, id string) (*model.Badge, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err)
	}

	badge, err := r.badgeService.Delete(int64(parsedID))

	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *queryResolver) Icons(_ context.Context) ([]*model.Icon, error) {
	icons, err := r.iconService.GetAll()

	if err != nil {
		return nil, err
	}

	return icons, nil
}

func (r *queryResolver) Icon(_ context.Context, id string) (*model.Icon, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err)
	}

	icon, err := r.iconService.GetOne(int64(parsedID))

	if err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *queryResolver) IconsByOwner(_ context.Context, _ int, _ string) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IconsByOwnerAndType(_ context.Context, _ int, _ string, _ string) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IconsByType(_ context.Context, _ string) ([]*model.Icon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Badges(_ context.Context) ([]*model.Badge, error) {
	badges, err := r.badgeService.GetAll()

	if err != nil {
		return nil, err
	}

	return badges, nil
}

func (r *queryResolver) Badge(_ context.Context, id string) (*model.Badge, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err)
	}

	badge, err := r.badgeService.GetOne(int64(parsedID))

	if err != nil {
		return nil, err
	}

	return badge, nil
}

func (r *queryResolver) BadgesByOwner(_ context.Context, _ int, _ string) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BadgesByOwnerAndType(_ context.Context, _ int, _ string, _ string) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BadgesByType(_ context.Context, _ string) ([]*model.Badge, error) {
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
