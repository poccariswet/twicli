package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func search(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "search")
		fmt.Println("\nInvalid arguments\nYou should $ twicli search 'sentence' ")
		return nil
	}
	text := c.Args()[0]
	searchResult, _ := api.GetSearch(text, nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println("[" + tweet.User.Name + "]" + " " + tweet.Text + "\n")
	}
	return nil
}
