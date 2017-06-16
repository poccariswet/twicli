package main

import (
	"fmt"
	"strconv"

	ui "github.com/gizak/termui"
	"github.com/urfave/cli"
)

func user(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "user")
		fmt.Println("\nInvalid arguments\nYou should $ twicli user 'username(@***)' ")
		return nil
	}
	username := c.Args()[0]
	user, _ := api.GetUsersShow(username, nil)

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	par0 := ui.NewPar("[User Information](fg-yellow)")
	par0.Height = 1
	par0.Width = 20
	par0.Y = 1
	par0.Border = false

	par1 := ui.NewPar(user.Name)
	par1.Height = 3
	par1.Width = len(user.Name) + 2
	par1.X = 20
	par1.BorderLabel = "Name"

	par2 := ui.NewPar(user.Description)
	par2.Height = 10
	if i := 22 + len(user.Name); i > 37 {
		par2.Width = i
	} else {
		par2.Width = 37
	}
	par2.Y = 4
	par2.BorderLabel = "Description"

	par3 := ui.NewPar("Follow: " + strconv.Itoa(user.FriendsCount) + "\nFollower: " + strconv.Itoa(user.FollowersCount))
	par3.Height = 4
	if i := 22 + len(user.Name); i > 37 {
		par3.Width = i
	} else {
		par3.Width = 37
	}
	par3.Y = 14
	par3.BorderLabel = "number of follows"

	ui.Render(par0, par1, par2, par3)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()

	return nil
}
