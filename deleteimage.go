package main

import (
	"fmt"
	"os"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
	"github.com/urfave/cli"
)

func deleteimage(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "deleteimage")
		return fmt.Errorf("\ninvalid arguments\nThere is too many words\n")
	}

	dirs := Dirplace(root)
	var files string
	for _, dir := range dirs {
		files += fmt.Sprintf("%s\n", dir)
	}

	imagename, err := pipeline.Output(
		[]string{"echo", files},    // pictures in the directory.
		[]string{os.Getenv(fuzzy)}, // fuzzy search
	)
	if err != nil {
		if strings.Contains("exit status 130", err.Error()) {
			return nil
		}
		return fmt.Errorf("cannot start fuzzy-search: %s", err)
	}

	if err := os.Remove(strings.TrimSpace(string(imagename))); err != nil { //delete file
		return fmt.Errorf("\ncannot delete image")
	}
  fmt.Println("delete success")
  return nil
}
