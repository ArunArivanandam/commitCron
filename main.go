package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v61/github"
)

func main() {
	fmt.Println("############ my repos ############")

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken("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")

	repos, _, err := client.Repositories.List(ctx, "ArunArivanandam", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		fmt.Println(*repo.FullName)
	}

	fmt.Println("############################")

	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	commits, _, err := client.Repositories.ListCommits(ctx, "ArunArivanandam", "greenlight", opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, commit := range commits {
		fmt.Printf("%s: %s\n", *commit.SHA, *commit.Commit.Message)
	}

}
