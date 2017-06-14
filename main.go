package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

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
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"CON_KEY_TW")
		os.Exit(1)
	}
	if ConSecKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"CON_SECRET_KEY_TW")
		os.Exit(1)
	}
	if AccKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"ACC_KEY_TW")
		os.Exit(1)
	}
	if AccSecKey == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"ACC_SECRET_KEY_TW")
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
			Usage:   "Tweet anything like this, $ twitter tweet '***' ",
			Action:  tweet,
		},

		{
			Name:    "search",
			Aliases: []string{"se", "s"},
			Usage:   "Search latest 15 tweets",
			Action:  search,
		},

		{
			Name:    "timeline",
			Aliases: []string{"ti", "time"},
			Usage:   "You can see your twitter timelines of 15",
			Action:  timeline,
		},

		{
			Name:    "usershow ",
			Aliases: []string{"sh", "u"},
			Usage:   "Show user's status or profile",
			Action:  usershow,
		},

		{
			Name:    "pictweet",
			Aliases: []string{"pic", "p"},
			Usage:   "Tweet picture and sentence like this, $ twitter pic 'imagename' 'sentence' ",
			Action:  pictweet,
		},

		{
			Name:    "saveimage",
			Aliases: []string{"sa"},
			Usage:   "Save image of url like thi, $ twitter sa 'url' 'imageName' ",
			Action:  saveimage,
		},
	}

	app.Run(os.Args)
}
