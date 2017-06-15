package main

import (
	"fmt"
	"net/url"

	"github.com/urfave/cli"
)

func timeline(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "timeline")
    fmt.Println("\nInvalid arguments\nYou should $ twicli timeline")
		return nil
	}

	v := url.Values{}
	timelines, _ := api.GetHomeTimeline(v) //default's numbers of timeline are 15
	for _, timeline := range timelines {
		fmt.Println("[" + timeline.User.Name + "]" + " " +timeline.Text)
	}
	return nil
}
