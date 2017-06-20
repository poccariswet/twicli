package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func delete(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "delete")
		fmt.Println("\nInvalid arguments\nYou should $ twicli delete")
		return nil
	}

	imagename, err := getimage()
	if err != nil {
		return err
	}

	if err := os.Remove(imagename); err != nil { //delete file
		return fmt.Errorf("\ncannot delete image")
	}
	fmt.Println("delete success")
	return nil
}
