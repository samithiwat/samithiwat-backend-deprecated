package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

type IconRepository struct {
	db database.Database
}

func NewIconRepository(db database.Database) *IconRepository {
	return &IconRepository{db: db}
}

func (r *IconRepository) FindAll(icons *[]*model.Icon) error {
	return r.db.GetConnection().Find(&icons).Error
}

func (r *IconRepository) FindOne(id int64, icon *model.Icon) error {
	return r.db.GetConnection().First(&icon, id).Error
}

func (r *IconRepository) Create(icon *model.Icon) error {
	return r.db.GetConnection().Create(&icon).Error
}

func (r *IconRepository) Update(id int64, icon *model.Icon) error {
	return r.db.GetConnection().Where(id).Updates(&icon).First(&icon).Error
}

func (r *IconRepository) Delete(id int64, icon *model.Icon) error {
	return r.db.GetConnection().First(&icon, id).Delete(&model.Icon{}).Error
}
