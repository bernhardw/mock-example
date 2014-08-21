package main

import (
	"fmt"
	"log"

	"github.com/ongoingio/mock/service"
)

// GetRepos fetches the repositories via the service (interface).
func GetRepos(s service.Servicer) []service.Repo {
	repos, err := s.Fetch()
	if err != nil {
		log.Fatal(err)
	}
	return repos
}

func main() {
	service := &service.Service{}
	repos := GetRepos(service)
	for _, repo := range repos {
		fmt.Printf("%s - %s\n", repo.Name, repo.Description)
	}
}
