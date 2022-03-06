package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllIcon(icons *[]*model.Icon) error {
	return r.db.GetConnection().Find(&icons).Error
}

func (r *GormRepository) FindIcon(id int64, icon *model.Icon) error {
	return r.db.GetConnection().First(&icon, id).Error
}

func (r *GormRepository) CreateIcon(icon *model.Icon) error {
	return r.db.GetConnection().Create(&icon).Error
}

func (r *GormRepository) UpdateIcon(id int64, icon *model.Icon) error {
	return r.db.GetConnection().Where(id).Updates(&icon).First(&icon).Error
}

func (r *GormRepository) DeleteIcon(id int64, icon *model.Icon) error {
	return r.db.GetConnection().First(&icon, id).Delete(&model.Icon{}).Error
}
