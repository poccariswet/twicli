package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func hello(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "hello")
		fmt.Println("\nInvalid arguments")
		return nil
	}
	fmt.Println("hello")
	return nil
}
