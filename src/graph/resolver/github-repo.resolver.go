package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (r *githubRepoResolver) DeletedAt(_ context.Context, obj *model.GithubRepo) (*time.Time, error) {
	if cmp.Equal(obj.DeletedAt, time.Time{}) {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not Found")
	}
	return nil, nil
}

func (r *mutationResolver) CreateGithubRepo(_ context.Context, newGithubRepo model.NewGithubRepo) (*model.GithubRepo, error) {
	githubRepo, err := r.githubRepoService.Create(&newGithubRepo)
	if err != nil {
		return nil, err
	}

	return githubRepo, nil
}

func (r *mutationResolver) UpdateGithubRepo(_ context.Context, id string, newGithubRepo model.NewGithubRepo) (*model.GithubRepo, error) {
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

func (r *mutationResolver) DeleteGithubRepo(_ context.Context, id string) (*model.GithubRepo, error) {
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

func (r *queryResolver) GithubRepos(_ context.Context) ([]*model.GithubRepo, error) {
	repo, err := r.githubRepoService.GetAll()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *queryResolver) GithubRepo(_ context.Context, id string) (*model.GithubRepo, error) {
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

// GithubRepo returns generated.GithubRepoResolver implementation.
func (r *Resolver) GithubRepo() generated.GithubRepoResolver { return &githubRepoResolver{r} }

type githubRepoResolver struct{ *Resolver }
