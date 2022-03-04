package graph

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"strconv"
	"time"
)

func (r *queryResolver) GithubRepos(ctx context.Context) ([]*model.GithubRepo, error) {
	repo, err := r.githubRepoService.GetAll()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *queryResolver) GithubRepo(ctx context.Context, id string) (*model.GithubRepo, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	githubRepo, err := r.githubRepoService.GetOne(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return githubRepo, nil
}

func (r *mutationResolver) CreateGithubRepo(ctx context.Context, newGithubRepo model.NewGithubRepo) (*model.GithubRepo, error) {
	githubRepo, err := r.githubRepoService.Create(&newGithubRepo)
	if err != nil {
		return nil, err
	}

	return githubRepo, nil
}

func (r *mutationResolver) UpdateGithubRepo(ctx context.Context, id string, newGithubRepo model.NewGithubRepo) (*model.GithubRepo, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	githubRepo, err := r.githubRepoService.Update(int64(parsedID), &newGithubRepo)
	if err != nil {
		return nil, err
	}

	return githubRepo, nil
}

func (r *mutationResolver) DeleteGithubRepo(ctx context.Context, id string) (*model.GithubRepo, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	githubRepo, err := r.githubRepoService.Delete(int64(parsedID))
	if err != nil {
		return nil, err
	}

	return githubRepo, nil
}

func (g githubRepoResolver) DeletedAt(ctx context.Context, obj *model.GithubRepo) (*time.Time, error) {
	//TODO implement me
	panic("implement me")
}


// GithubRepo returns generated.GithubRepoResolver implementation.
func (r *Resolver) GithubRepo() generated.GithubRepoResolver { return &githubRepoResolver{r} }

type githubRepoResolver struct{ *Resolver }
