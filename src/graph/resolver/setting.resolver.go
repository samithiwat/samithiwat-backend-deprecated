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

func (r *aboutMeResolver) DeletedAt(_ context.Context, _ *model.AboutMe) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSetting(_ context.Context, newSetting model.NewSetting) (*model.Setting, error) {
	setting, err := r.settingService.Create(&newSetting)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) UpdateSetting(_ context.Context, id string, newSetting *model.NewSetting) (*model.Setting, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.settingService.Update(int64(parsedID), newSetting)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) DeleteSetting(_ context.Context, id string) (*model.Setting, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.settingService.Delete(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) CreateAboutMe(_ context.Context, newAboutMe model.NewAboutMe) (*model.AboutMe, error) {
	setting, err := r.aboutMeSettingService.Create(&newAboutMe)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) UpdateAboutMe(_ context.Context, id string, newAboutMe *model.NewAboutMe) (*model.AboutMe, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.aboutMeSettingService.Update(int64(parsedID), newAboutMe)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) DeleteAboutMe(_ context.Context, id string) (*model.AboutMe, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.aboutMeSettingService.Delete(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) CreateTimeline(_ context.Context, newTimeline model.NewTimeline) (*model.Timeline, error) {
	setting, err := r.timelineSettingService.Create(&newTimeline)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) UpdateTimeline(_ context.Context, id string, newTimeline *model.NewTimeline) (*model.Timeline, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.timelineSettingService.Update(int64(parsedID), newTimeline)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *mutationResolver) DeleteTimeline(_ context.Context, id string) (*model.Timeline, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.timelineSettingService.Delete(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *queryResolver) Settings(_ context.Context) ([]*model.Setting, error) {
	setting, err := r.settingService.GetAll()
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *queryResolver) Setting(_ context.Context, id string) (*model.Setting, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.settingService.GetOne(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *queryResolver) AboutMes(_ context.Context) ([]*model.AboutMe, error) {
	setting, err := r.aboutMeSettingService.GetAll()
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *queryResolver) AboutMe(_ context.Context, id string) (*model.AboutMe, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.aboutMeSettingService.GetOne(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *queryResolver) Timelines(_ context.Context) ([]*model.Timeline, error) {
	setting, err := r.timelineSettingService.GetAll()
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *queryResolver) Timeline(_ context.Context, id string) (*model.Timeline, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	setting, err := r.timelineSettingService.GetOne(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *settingResolver) DeletedAt(_ context.Context, _ *model.Setting) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *timelineResolver) SettingID(_ context.Context, _ *model.Timeline) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *timelineResolver) DeletedAt(_ context.Context, _ *model.Timeline) (*time.Time, error) {
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
