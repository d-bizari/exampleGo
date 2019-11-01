package main

import (
	"github.com/abiosoft/ishell"
	"github.com/d-bizari/exampleGo/src/domain"
	"github.com/d-bizari/exampleGo/src/service"
	"strconv"
)

func main() {
	service.InitializeService()
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

			id, err := service.PublishTweet(tweet)

			if err != nil {
				println(err.Error())
			} else {
				c.Println(id)
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
				c.Println("Tweet:", tweets[i].Text)
				c.Println("User:", tweets[i].User)
				c.Println()
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts the tweets of the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write user: ")
			user := c.ReadLine()

			count := service.CountTweetsByUser(user)
			c.Println(count)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetByID",
		Help: "Shows tweet by ID",
		Func: func(c *ishell.Context) {
			var id int
			defer c.ShowPrompt(true)

			c.Print("Write ID: ")
			idStr := c.ReadLine()

			id, _ = strconv.Atoi(idStr)

			tweet := service.GetTweetById(int64(id))

			c.Println("Tweet:", tweet.User)
			c.Println("User:", tweet.Text)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUser",
		Help: "Shows tweets by user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write User: ")
			user := c.ReadLine()

			tweets := service.GetTweetsByUser(user)

			for _ , tweet := range tweets {
				c.Println("Tweet:", tweet.User)
				c.Println("User:", tweet.Text)
			}

			return
		},
	})

	shell.Run()

}
