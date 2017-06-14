package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func tweet(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "tweet")
		return fmt.Errorf("\ninvalid arguments words or you forgot ' '")
	}
	text := c.Args()[0]
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		return fmt.Errorf("\nfailed tweet")
	}

	if tweet.Text != "" {
		fmt.Println(tweet.Text)
		fmt.Println("success!")
	}
	return nil
}