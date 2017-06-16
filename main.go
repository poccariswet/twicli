package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

const fuzzy = "FZF"

var (
	ConKey    = os.Getenv("CON_KEY_TW")        //Twitter-Consumer-Key
	ConSecKey = os.Getenv("CON_SECRET_KEY_TW") //Twitter-Consumer-Secret-Key
	AccKey    = os.Getenv("ACC_KEY_TW")        //Twitter-Access-Key
	AccSecKey = os.Getenv("ACC_SECRET_KEY_TW") //Twitter-Access-Secret-Key
	api       *anaconda.TwitterApi

	path = os.Getenv("HOME")
	root = path + "/.images"
)

func Set() {
	anaconda.SetConsumerKey(ConKey)
	anaconda.SetConsumerSecret(ConSecKey)
	api = anaconda.NewTwitterApi(AccKey, AccSecKey)
	// fmt.Println(api)
}

func init() {
	Set() //keyの設定
	if ConKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter consmer key")
		os.Exit(1)
	}
	if ConSecKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter secret consumer key")
		os.Exit(1)
	}
	if AccKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter access token")
		os.Exit(1)
	}
	if AccSecKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter access secret key")
		os.Exit(1)
	}

	_, err := os.Stat(root)
	if err != nil {
		os.Mkdir(root, 0777)
		fmt.Println("Made a Dir at" + root)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Twitter CLI"
	app.Usage = "Command line interface for Twitter"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "tweet",
			Aliases: []string{"tw", "t"},
			Usage:   "Tweet sentence",
			Action:  tweet,
		},

		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search latest 15 tweets",
			Action:  search,
		},

		{
			Name:    "timeline",
			Aliases: []string{"ti", "time"},
			Usage:   "Show twitter timelines of 15",
			Action:  timeline,
		},

		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Show user's profile",
			Action:  user,
		},

		{
			Name:    "pictweet",
			Aliases: []string{"pic", "p"},
			Usage:   "Tweet picture and sentence",
			Action:  pictweet,
		},

		{
			Name:    "save",
			Aliases: []string{"sa"},
			Usage:   "Save image",
			Action:  save,
		},

		{
			Name: "delete",
			Aliases: []string{"de", "d"},
			Usage: "Delete in the HOME/.images image",
			Action: delete,
		},
	}

	app.Run(os.Args)
}
