package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func usershow(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "usershow")
		return fmt.Errorf("\ninvalid arguments\nYou should $ user 'username(@***)' ")
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
