package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllImage(images *[]*model.Image) error {
	return r.db.GetConnection().Find(&images).Error
}

func (r *GormRepository) FindImage(id int64, image *model.Image) error {
	return r.db.GetConnection().First(&image, id).Error
}

func (r *GormRepository) CreateImage(image *model.Image) error {
	return r.db.GetConnection().Create(&image).Error
}

func (r *GormRepository) UpdateImage(id int64, image *model.Image) error {
	return r.db.GetConnection().Where(id).Updates(&image).First(&image).Error
}

func (r *GormRepository) DeleteImage(id int64, image *model.Image) error {
	return r.db.GetConnection().First(&image, id).Delete(&model.Image{}).Error
}
