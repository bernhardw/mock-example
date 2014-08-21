package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://api.github.com/orgs/ongoingio/repos"

// Servicer defines an interface to be implemented by objects to fetch repositories.
type Servicer interface {
	Fetch() ([]Repo, error)
}

// Service serves as a service type.
type Service struct{}

// The Repo type represents a single repository on Github.
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Fetch retrieves all repositories from an organization and returns them as a slice of Repo.
func (service *Service) Fetch() ([]Repo, error) {
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
