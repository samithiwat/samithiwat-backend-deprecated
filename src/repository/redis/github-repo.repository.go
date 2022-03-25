package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com"
const GITHUB_USER = "samithiwat"

func (r *RedisRepository) GetGithubRepoDetails() error {
	res, err := http.Get(fmt.Sprintf("%s/users/%s/repos", GITHUB_API_URL, GITHUB_USER))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Printf("Body %s", body)

	return nil
}
