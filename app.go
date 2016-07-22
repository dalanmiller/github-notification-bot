package main

import (
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func getNotifications() {
	SINCE := new(time.Time)

	// Setup Github Auth
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List("", nil)
	opts := github.NotificationListOptions{}
	if SINCE != nil {
		opts.Since = *SINCE
	}

	opts.All = false
	opts.Participating = true

	notifications, response, err := client.Activity.ListNotifications(&opts)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(notifications)
	spew.Dump(response)
}

func main() {
	getNotifications()
	// api := slack.New(os.Getenv("SLACK_TOKEN"))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	// groups, err := api.GetGroups(false)
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	return
	// }
	// for _, group := range groups {
	// 	fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	// }

}
