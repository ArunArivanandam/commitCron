package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v61/github"
)

func listAllCommits(client *github.Client, ctx context.Context, owner string, opts *github.CommitsListOptions) {
	commitsAll := make(map[string][]string)
	
	repos, err := getReposList(client, ctx, owner)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, repo := range repos {
		commits, _ := getCommitsList(client, ctx, owner, repo, opts)
		commitsAll[repo] = commits
	}

	for repo, commits := range commitsAll {
		fmt.Printf("%s: %s\n", repo, commits)
	}
	
	
}