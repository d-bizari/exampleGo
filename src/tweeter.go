package main

import (
	"github.com/abiosoft/ishell"
	"github.com/d-bizari/exampleGo/src/domain"
	"github.com/d-bizari/exampleGo/src/service"
	"strconv"
)

func main() {
	tm := service.NewTweetManager()

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

			tweet := domain.NewTextTweet(user, texto)

			id, err := tm.PublishTweet(tweet)

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

			tweet := tm.GetTweet()

			c.Println("Tweet:", tweet.GetUser())
			c.Println("User:", tweet.GetText())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tm.GetTweets()

			for i := 0; i < len(tweets); i++ {
				c.Println("Tweet:", tweets[i].GetText())
				c.Println("User:", tweets[i].GetUser())
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

			count := tm.CountTweetsByUser(user)
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

			tweet := tm.GetTweetById(int64(id))

			c.Println("Tweet:", tweet.GetUser())
			c.Println("User:", tweet.GetText())

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

			tweets := tm.GetTweetsByUser(user)

			for _, tweet := range tweets {
				c.Println("Tweet:", tweet.GetUser())
				c.Println("User:", tweet.GetText())
			}

			return
		},
	})

	shell.Run()

}
