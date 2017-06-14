package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func search(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "search")
		return fmt.Errorf("\ninvalid arguments of multipl words or you forgot to add at sentence ' ' ")
	}
	text := c.Args()[0]
	searchResult, _ := api.GetSearch(text, nil)
	for _, tweet := range searchResult.Statuses {
		// fmt.Printf("%d  :", i+1)
		fmt.Println("[" + tweet.User.Name + "]" + " " + tweet.Text + "\n")
	}
	return nil
}