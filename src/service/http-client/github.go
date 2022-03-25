package httpClient

import (
	"encoding/json"
	"github.com/samithiwat/samithiwat-backend/src/dto"
)

const API_URL = ""

type GithubClient interface {
	getAllRepos() *dto.GithubRepo
}

type Github struct {
	Client Client
}

func NewGithubClient(client Client) *Github {
	return &Github{
		Client: client,
	}
}

func (g *Github) getAllRepos() (*dto.GithubRepo, error) {
	var repos dto.GithubRepo

	body, err := g.Client.GET("https://api.github.com")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*body, &repos)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
