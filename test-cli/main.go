package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "test CLI"
	app.Usage = "Command line interface for test"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"h"},
			Usage:   "print out test",
			Action:  hello,
		},
	}
	app.Run(os.Args)
}
