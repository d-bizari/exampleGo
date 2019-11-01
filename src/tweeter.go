package main

import (
	"github.com/abiosoft/ishell"
	"github.com/d-bizari/exampleGo/src/domain"
	"github.com/d-bizari/exampleGo/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")
			texto := c.ReadLine()

			c.Print("Write your user: ")
			user := c.ReadLine()

			tweet := domain.NewTweet(texto, user)

			err := service.PublishTweet(tweet)

			if err != nil {
				println(err.Error())
			} else {
				c.Print("Tweet sent\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println("Tweet:", tweet.User)
			c.Println("User:", tweet.Text)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for i := 0; i < len(tweets); i++ {
				c.Println("Tweet:", tweets[i].User)
				c.Println("User:", tweets[i].Text)
				c.Println()
			}

			return
		},
	})

	shell.Run()

}
