package main

import (
	"fmt"
	"log"

	"github.com/ongoingio/mock/service"
)

// GetRepos gets the repos from the service.
func GetRepos() []service.Repo {
	repos, err := service.Fetch()
	if err != nil {
		log.Fatal(err)
	}
	return repos
}

func main() {
	repos := GetRepos()
	for _, repo := range repos {
		fmt.Printf("%s - %s\n", repo.Name, repo.Description)
	}
}
