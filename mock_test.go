package main

import (
	"testing"

	"github.com/ongoingio/mock/service"
)

// We create our own custom service mock to be used in place of the service that queries the Github API.
type serviceMock struct{}

func (s *serviceMock) Fetch() ([]service.Repo, error) {
	repos := []service.Repo{
		service.Repo{"Foo", "Foo desc."},
	}
	return repos, nil
}

func TestGetRepos(t *testing.T) {
	service := &serviceMock{}
	repos := GetRepos(service)
	if repos[0].Name != "Foo" {
		t.Fatalf("Name should be `Foo`, is `%s`", repos[0].Name)
	}
}
