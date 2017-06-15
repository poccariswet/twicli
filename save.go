package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func save(c *cli.Context) error {
	if c.NArg() != 2 {
		cli.ShowCommandHelp(c, "save")
		fmt.Println("\nInvalid arguments\nYou should $ twicli save 'url' 'image name' ")
		return nil
	}

	url := c.Args()[0]
	imageName := c.Args()[1]

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("This URL is not validity.")
	}
	defer res.Body.Close()

	picture, err := os.Create(fmt.Sprintf(root+"/%s.jpg", imageName))
	if err != nil {
		return fmt.Errorf("failed create image.")
	} else {
		fmt.Println("success")
	}

	defer picture.Close()
	io.Copy(picture, res.Body)

	return nil
}
