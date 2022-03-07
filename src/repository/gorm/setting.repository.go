package repository

import (
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllSetting(setting *[]*model.Setting) error {
	return r.db.GetConnection().Find(&setting).Error
}

func (r *GormRepository) FindSetting(id int64, setting *model.Setting) error {
	return r.db.GetConnection().Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").First(&setting, id).Error
}

func (r *GormRepository) FindActiveSetting(setting *model.Setting) error {
	return r.db.GetConnection().Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").Where("isActivated = ?", true).Take(&setting).Error
}

func (r *GormRepository) CreateSetting(setting *model.Setting) error {
	return r.db.GetConnection().Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").Create(&setting).Error
}

func (r *GormRepository) UpdateSetting(id int64, setting *model.Setting) error {
	err := r.db.GetConnection().Where(id).Updates(&setting).First(&setting).Error

	if (!cmp.Equal(setting.Timeline, model.Timeline{})) {
		err = r.db.GetConnection().Model(&setting).Association("Timeline").Replace(&setting.Timeline)
		err = r.db.GetConnection().Model(&setting.Timeline).Association("Icon").Replace(&setting.Timeline.Icon)
		err = r.db.GetConnection().Model(&setting.Timeline).Association("Images").Replace(&setting.Timeline.Images)
	}

	if (!cmp.Equal(setting.AboutMe, model.AboutMe{})) {
		err = r.db.GetConnection().Model(&setting).Association("AboutMe").Replace(&setting.AboutMe)
	}

	return err
}

func (r *GormRepository) DeleteSetting(id int64, setting *model.Setting) error {
	return r.db.GetConnection().First(&setting, id).Delete(&model.Setting{}).Error
}
