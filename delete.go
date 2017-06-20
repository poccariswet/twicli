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

	// dirs := Dirplace(root)
	// var files string
	// for _, dir := range dirs {
	// 	files += fmt.Sprintf("%s\n", dir)
	// }
	//
	// imagename, err := pipeline.Output(
	// 	[]string{"echo", files},    // pictures in the directory.
	// 	[]string{os.Getenv(fuzzy)}, // fuzzy search
	// )
	// if err != nil {
	// 	if strings.Contains("exit status 130", err.Error()) {
	// 		return nil
	// 	}
	// 	return fmt.Errorf("cannot start fuzzy-search: %s", err)
	// }
	//
	// if strings.TrimSpace(string(imagename)) == "" {
	// 	fmt.Println("cannnot select empty string\n")
	// 	return fmt.Errorf("Please select image")
	// }
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
