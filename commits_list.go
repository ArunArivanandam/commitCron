package main

import (
	"context"
	"log"

	"github.com/google/go-github/v61/github"
)

func getCommitsList(client *github.Client, ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]string, error) {

	commits_list := []string{}

	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, commit := range commits {
		commits_list = append(commits_list, *commit.Commit.Message)
	}
	return commits_list, nil
}