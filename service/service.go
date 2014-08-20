package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://api.github.com/orgs/ongoingio/repos"

// The Repo type represents a single repository on Github.
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Fetch retrieves all repositories from an organization and returns them as a slice of Repo.
func Fetch() ([]Repo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var repos []Repo
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
