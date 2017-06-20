package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"

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

func pictweet(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "pictweet")
		fmt.Println("\nInvalid arguments\nYou should$ twicli pic 'sentence'")
		return nil
	}

	dirs := Dirplace(root)
	var files string
	for _, dir := range dirs {
		files += fmt.Sprintf("%s\n", dir)
	}

	imagename, err := getimage()
	if err != nil {
		return err
	}
	text := c.Args()[0]
	pic := encode(imagename)
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
