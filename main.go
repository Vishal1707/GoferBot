package main

import (
	"fmt"
	"os"

	fb "github.com/huandu/facebook"
	"github.com/urfave/cli"
)

const (
	appname     = "FacebookBot"
	appver      = "0.0.1"
	accessToken = "access_token"
)

var (
	key      string
	response string
)

type Posts struct {
	Data []Post
}

type Post struct {
	Message   string `json:"message"`
	ID        string `json:"id"`
	Story     string `json:"story"`
	Timestamp string `json:"created_time"`
}

func main() {
	app := cli.NewApp()
	app.Name = appname
	app.Usage = "Bot to respond to the facebook post."
	app.Version = appver
	app.Authors = []cli.Author{
		{
			Name:  "Vishal Kumar Singh",
			Email: "vishalkumarsingh1707@gmail.com",
		},
	}
	app.Copyright = "(c) 2017 Vishal Kumar Singh"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key",
			Value:       "Happy,Birthday",
			Usage:       "Key word to be searched in the post",
			EnvVar:      "FB_KEY",
			Destination: &key,
		},
		cli.StringFlag{
			Name:        "response,r",
			Value:       "Thank you",
			Usage:       "Response to be given",
			EnvVar:      "FB_RESPONSE",
			Destination: &response,
		},
	}
	app.Action = func(c *cli.Context) error {

		res, err := fb.Get("/me/feed", fb.Params{
			"access_token": accessToken,
		})

		if err != nil {
			fmt.Println("Error")
			return err
		}

		var posts Posts
		res.Decode(&posts)
		//fmt.Println(posts)

		for index, element := range posts.Data {
			// index is the index where we are
			// element is the element from someSlice for where we are
			if index == 1 {
				url := element.ID + "/likes"
				fmt.Println(url)
				fb.Post(url, fb.Params{
					"access_token": accessToken,
					//"message":      "I am a bot",
				})
			}
		}
		return nil
	}

	app.Run(os.Args)
}
