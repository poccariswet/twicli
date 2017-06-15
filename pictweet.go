package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
	"github.com/urfave/cli"
)

func encode(image string) string {
	file, _ := os.Open(image)
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

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

func pictweet(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "pictweet")
		return fmt.Errorf("\ninvalid arguments or multipl words")
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
	fmt.Println(strings.TrimSpace(string(imagename)))
	text := c.Args()[0]
	pic := encode(strings.TrimSpace(string(imagename)))
	media, _ := api.UploadMedia(pic)

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString) // 画像のidを付加してツイートする
	tweet, _ := api.PostTweet(text, v)
	if tweet.Text == "" {
		return fmt.Errorf("failed the tweet.\n")
	}

	fmt.Println(tweet.Text)
	fmt.Println("success!")

	return nil
}
