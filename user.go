package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func user(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "user")
		fmt.Println("\nInvalid arguments\nYou should $ twicli user 'username(@***)' ")
		return nil
	}
	text := c.Args()[0]
	user, _ := api.GetUsersShow(text, nil)
	fmt.Println(user.Name)
	fmt.Printf("Follow   :%d\n", user.FriendsCount)
	fmt.Printf("Follower :%d\n", user.FollowersCount)
	fmt.Printf("userID   :%d\n", user.Id)
	fmt.Println(user.Description)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println(user.ProfileImageURL)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println(user.ProfileBannerURL)
	fmt.Println("--------------------------------------------------------------")
	if user.URL == "" {
		fmt.Println("User dosen't have another URL.")
	} else {
		fmt.Println(user.URL)
	}
	return nil
}
