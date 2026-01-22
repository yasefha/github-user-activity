package main

import (
	"fmt"
	"github-user-activity/cli"
	"github-user-activity/internal/github"
	"net/http"
	"os"
)

func main() {
	username, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	body, status, err := github.FetchActivity(username)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch status {
	case http.StatusOK:
		events, err := github.ParseActivity(body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%s activity: \n", username)
		for i, event := range events {
			fmt.Println(i+1, "-", event)
		}

	case http.StatusNotFound:
		fmt.Fprintln(os.Stderr, "user not found")
		os.Exit(1)

	case http.StatusForbidden:
		fmt.Fprintln(os.Stderr, "rate limit reached")
		os.Exit(1)

	default:
		fmt.Fprintf(os.Stderr, "unexpected response: %d\n", status)
		os.Exit(1)
	}

}
