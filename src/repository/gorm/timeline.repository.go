package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllTimeline(timelines *[]*model.Timeline) error {
	return r.db.GetConnection().Find(timelines).Error
}

func (r *GormRepository) FindOneTimeline(id int64, timeline *model.Timeline) error {
	return r.db.GetConnection().Preload("Images").Preload("Icon").First(timeline, id).Error
}

func (r *GormRepository) CreateTimeline(timeline *model.Timeline) error {
	return r.db.GetConnection().Preload("Images").Preload("Icon").Create(timeline).Error
}

func (r *GormRepository) UpdateTimeline(id int64, timeline *model.Timeline) error {
	err := r.db.GetConnection().Where(id).Updates(timeline).First(timeline).Error

	if (timeline.Icon != model.Icon{}) {
		err = r.db.GetConnection().Model(&timeline).Association("Icon").Replace(&timeline.Icon)
	}

	if len(timeline.Images) > 0 {
		err = r.db.GetConnection().Model(&timeline).Association("Images").Replace(timeline.Images)
	}

	return err
}

func (r *GormRepository) DeleteTimeline(id int64, timeline *model.Timeline) error {
	return r.db.GetConnection().First(&timeline, id).Delete(&model.Timeline{}).Error
}
