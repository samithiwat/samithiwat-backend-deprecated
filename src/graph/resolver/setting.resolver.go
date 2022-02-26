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

func (r *aboutMeResolver) DeletedAt(ctx context.Context, obj *model.AboutMe) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSetting(ctx context.Context, newSetting model.NewSetting) (*model.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSetting(ctx context.Context, id string, newSetting model.NewSetting) (*model.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSetting(ctx context.Context, id string) (*model.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAboutMe(ctx context.Context, name string, description string, content string, imgURL string, settingID string) (*model.AboutMe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAboutMe(ctx context.Context, id string, name string, description string, content string, imgURL string, settingID string) (*model.AboutMe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAboutMe(ctx context.Context, id string) (*model.AboutMe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTimeline(ctx context.Context, slug string, name string, description string, thumbnail string, eventDate time.Time, iconID string, settingID string) (*model.Timeline, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTimeline(ctx context.Context, id string, slug string, name string, description string, thumbnail string, eventDate time.Time, iconID string, settingID string) (*model.Timeline, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTimeline(ctx context.Context, id string) (*model.Timeline, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Settings(ctx context.Context) ([]*model.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Setting(ctx context.Context, id string) (*model.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AboutMes(ctx context.Context) ([]*model.AboutMe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AboutMe(ctx context.Context, id string) (*model.AboutMe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timelines(ctx context.Context) ([]*model.Timeline, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timeline(ctx context.Context, id string) (*model.Timeline, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *settingResolver) AboutMe(ctx context.Context, obj *model.Setting) (*model.AboutMe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *settingResolver) Timeline(ctx context.Context, obj *model.Setting) (*model.Timeline, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *settingResolver) DeletedAt(ctx context.Context, obj *model.Setting) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *timelineResolver) DeletedAt(ctx context.Context, obj *model.Timeline) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// AboutMe returns generated.AboutMeResolver implementation.
func (r *Resolver) AboutMe() generated.AboutMeResolver { return &aboutMeResolver{r} }

// Setting returns generated.SettingResolver implementation.
func (r *Resolver) Setting() generated.SettingResolver { return &settingResolver{r} }

// Timeline returns generated.TimelineResolver implementation.
func (r *Resolver) Timeline() generated.TimelineResolver { return &timelineResolver{r} }

type aboutMeResolver struct{ *Resolver }
type settingResolver struct{ *Resolver }
type timelineResolver struct{ *Resolver }
