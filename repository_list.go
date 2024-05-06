package main

import (
	"context"
	"log"

	"github.com/google/go-github/v61/github"
)

func getReposList(client *github.Client, ctx context.Context, owner string) ([]string, error) {

	repo_list := []string{}

	repos, _, err := client.Repositories.ListByUser(ctx, owner, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		repo_list = append(repo_list, *repo.Name)
	}
	return repo_list, nil
}