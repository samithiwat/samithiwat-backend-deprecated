package repository

import (
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllBadge(badges *[]*model.Badge) error {
	return r.db.GetConnection().Find(&badges).Error
}

func (r *GormRepository) FindOneBadge(id int64, badge *model.Badge) error {
	return r.db.GetConnection().Preload("Icon").First(&badge, id).Error
}

func (r *GormRepository) CreateBadge(badge *model.Badge) error {
	return r.db.GetConnection().Create(&badge).Error
}

func (r *GormRepository) UpdateBadge(id int64, badge *model.Badge) error {
	err := r.db.GetConnection().Where(id).Updates(&badge).First(&badge).Error

	if err != nil {
		return err
	}

	if (!cmp.Equal(badge.Icon, model.Icon{})) {
		err = r.db.GetConnection().Model(&badge).Association("Icon").Replace(&badge.Icon)
	}

	return err
}

func (r *GormRepository) DeleteBadge(id int64, badge *model.Badge) error {
	return r.db.GetConnection().First(&badge, id).Delete(&model.Badge{}).Error
}
