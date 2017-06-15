package main

import (
	"fmt"
	"net/url"

	"github.com/urfave/cli"
)

func timeline(c *cli.Context) error {
	v := url.Values{}
	timelines, _ := api.GetHomeTimeline(v) //default's numbers of timeline are 15
	for _, timeline := range timelines {
		// fmt.Printf("%3d:", i+1)
		fmt.Println("[" + timeline.User.Name + "]" + " " +timeline.Text)
	}
	return nil
}
