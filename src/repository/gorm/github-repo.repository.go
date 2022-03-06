package repository

import (
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *GormRepository) FindAllGithubRepo(repos *[]*model.GithubRepo) error {
	return r.db.GetConnection().Find(&repos).Error
}

func (r *GormRepository) FindGithubRepo(id int64, repo *model.GithubRepo) error {
	return r.db.GetConnection().Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").First(&repo, id).Error
}

func (r *GormRepository) CreateGithubRepo(repo *model.GithubRepo) error {
	return r.db.GetConnection().Preload("Language").Preload("Framework").Preload("Language.Icon").Preload("Framework.Icon").Create(&repo).Error
}

func (r *GormRepository) UpdateGithubRepo(id int64, repo *model.GithubRepo) error {
	err := r.db.GetConnection().Where(id).Updates(&repo).First(&repo).Error

	if err != nil {
		return err
	}

	if (!cmp.Equal(repo.Framework, model.Badge{})) {
		err = r.db.GetConnection().Model(&repo).Association("Framework").Replace(&repo.Framework)
		err = r.db.GetConnection().Model(&repo.Framework).Association("Icon").Append(&repo.Framework.Icon)
	}

	if (!cmp.Equal(repo.Language, model.Badge{})) {
		err = r.db.GetConnection().Model(&repo).Association("Language").Replace(&repo.Language)
		err = r.db.GetConnection().Model(&repo.Language).Association("Icon").Append(&repo.Language.Icon)
	}

	return err
}

func (r *GormRepository) DeleteGithubRepo(id int64, repo *model.GithubRepo) error {
	return r.db.GetConnection().First(&repo, id).Delete(&model.GithubRepo{}).Error
}
