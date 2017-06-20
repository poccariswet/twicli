package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
)

func Dirplace(dir string) []string { //finding in the picture file
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirplace(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func getimage() (string, error) {
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
			return "failed", fmt.Errorf("exit status 130")
		}
		return "failed", fmt.Errorf("cannot start fuzzy-search: %s", err)
	}

	if strings.TrimSpace(string(imagename)) == "" {
		fmt.Println("cannnot select empty string\n")
		return "failed", fmt.Errorf("Please select image")
	}

	return strings.TrimSpace(string(imagename)), nil
}
