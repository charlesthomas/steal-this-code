package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v63/github"
)

func newClient(token string) *client {
	return &client{
		gh:        github.NewClient(nil).WithAuthToken(token),
		remaining: 1, // don't know how the real value, but make sure we try the very first time
	}
}

type client struct {
	gh        *github.Client
	remaining int
	reset     time.Time
}

func (c *client) update(rate github.Rate) {
	c.remaining = rate.Remaining
	c.reset = rate.Reset.Time
}

func (c *client) wait() {
	if c.remaining == 0 {
		time.Sleep(time.Until(c.reset))
	}
}

func (c *client) ListRepos(owner string) ([]*github.Repository, error) {
	var page int
	var res []*github.Repository

	opts := new(github.RepositoryListByUserOptions)

	for {
		// here's the magic to wait until the rate limit is over
		c.wait()

		// make the API call
		opts.ListOptions.Page = page
		data, meta, err := c.gh.Repositories.ListByUser(context.Background(), owner, opts)
		if err != nil {
			return nil, err
		}

		// update the rate limit info for the next wait()
		c.update(meta.Rate)

		res = append(res, data...)
		if meta.NextPage <= page {
			break
		}
		page = meta.NextPage
	}

	return res, nil
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("export GITHUB_TOKEN=<some_token>")
		os.Exit(1)
	}

	client := newClient(token)
	repos, err := client.ListRepos("charlesthomas")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}
