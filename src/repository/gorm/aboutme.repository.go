package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllAboutMe(aboutMes *[]*model.AboutMe) error {
	return r.db.GetConnection().Find(&aboutMes).Error
}

func (r *GormRepository) FindAboutMe(id int64, aboutMe *model.AboutMe) error {
	return r.db.GetConnection().First(&aboutMe, id).Error
}

func (r *GormRepository) CreateAboutMe(aboutMe *model.AboutMe) error {
	return r.db.GetConnection().Create(&aboutMe).Error
}

func (r *GormRepository) UpdateAboutMe(id int64, aboutMe *model.AboutMe) error {
	return r.db.GetConnection().Where(id).Updates(&aboutMe).First(&aboutMe).Error
}

func (r *GormRepository) DeleteAboutMe(id int64, aboutMe *model.AboutMe) error {
	return r.db.GetConnection().First(&aboutMe, id).Delete(&model.AboutMe{}).Error
}
