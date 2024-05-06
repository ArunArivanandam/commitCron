package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/go-github/v61/github"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	loc, _ := time.LoadLocation("Asia/Kolkata")
    now := time.Now().In(loc)
    fmt.Println("Location : ", loc, " Time : ", now)

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(os.Getenv("ACCESS_KEY"))
	owner := "ArunArivanandam"


	date := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour)
	untilDate := time.Now().Truncate(24 * time.Hour)

	opts := &github.CommitsListOptions{
		Since: date,
		Until: untilDate,
	}


	s := gocron.NewScheduler(loc)

	s.Every(1).Day().At("20:00").Do(listAllCommits, client, ctx, owner, opts)
	
	s.StartBlocking()
}


