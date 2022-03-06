package graph

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
	"strconv"
	"time"
)

func (r *queryResolver) GithubRepos(_ context.Context) ([]*model2.GithubRepo, error) {
	repo, err := r.githubRepoService.GetAll()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *queryResolver) GithubRepo(_ context.Context, id string) (*model2.GithubRepo, error) {
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

func (r *mutationResolver) CreateGithubRepo(_ context.Context, newGithubRepo model2.NewGithubRepo) (*model2.GithubRepo, error) {
	githubRepo, err := r.githubRepoService.Create(&newGithubRepo)
	if err != nil {
		return nil, err
	}

	return githubRepo, nil
}

func (r *mutationResolver) UpdateGithubRepo(_ context.Context, id string, newGithubRepo model2.NewGithubRepo) (*model2.GithubRepo, error) {
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

func (r *mutationResolver) DeleteGithubRepo(_ context.Context, id string) (*model2.GithubRepo, error) {
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

func (g githubRepoResolver) DeletedAt(_ context.Context, obj *model2.GithubRepo) (*time.Time, error) {
	if cmp.Equal(obj.DeletedAt, time.Time{}) {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not Found")
	}
	return nil, nil
}

// GithubRepo returns generated.GithubRepoResolver implementation.
func (r *Resolver) GithubRepo() generated.GithubRepoResolver { return &githubRepoResolver{r} }

type githubRepoResolver struct{ *Resolver }
