package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
	"golang.org/x/oauth2"
	"os"
)

type labelslice []string

func (s *labelslice) String() string {
	return fmt.Sprintf("%v", labelargs)
}

func (s *labelslice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var labelargs labelslice

func main() {
	var (
		user  = flag.String("user", "", "string flag")
		body  = flag.String("body", "", "string flag")
		title = flag.String("title", "No Title", "string flag")
		repo  = flag.String("repo", "memo", "string flag")
	)
	flag.Var(&labelargs, "labels", "Data values.")
	flag.Parse()

	labels := make([]string, len(labelargs))
	copy(labels, labelargs)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	opt := &github.IssueRequest{
		Title:    title,
		Body:     body,
		Assignee: user,
		Labels:   &labels,
	}

	issue, _, err := client.Issues.Create(oauth2.NoContext, *user, *repo, opt)
	// pp.Print(client.Issues)
	// pp.Print(opt)
	// pp.Print(issue)
	if err != nil {
		pp.Print(err)
		return
	}
	fmt.Printf("https://github.com/%v/%v/issues/%v", *user, *repo, *issue.Number)
}
