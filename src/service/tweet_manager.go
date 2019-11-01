package service

import (
	"fmt"
	"github.com/d-bizari/exampleGo/src/domain"
)

var Tweets []*domain.Tweet
var Users map[string][]*domain.Tweet
var IdCounter int64

func PublishTweet(tweet *domain.Tweet) (int64, error) {
	if tweet.User == "" {
		return -1, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return -1, fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return -1, fmt.Errorf("characters exceeded, only 140 characters are allowed")
	}
	IdCounter++
	tweet.Id = IdCounter
	Tweets = append(Tweets, tweet)
	if Users[tweet.User] == nil {
		Users[tweet.User] = make([]*domain.Tweet, 0)
	}
	Users[tweet.User] = append(Users[tweet.User], tweet)
	return IdCounter, nil
}

func InitializeService() {
	Tweets = make([]*domain.Tweet, 0)
	IdCounter = 0
	Users = make(map[string][]*domain.Tweet)
}

func GetTweets() []*domain.Tweet {
	return Tweets
}

func GetTweet() *domain.Tweet {
	return Tweets[len(Tweets)-1]
}

func GetTweetById(id int64) *domain.Tweet {
	for _, tweet := range Tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func CountTweetsByUser(user string) int {
	count := 0

	for _, tweet := range Tweets {
		if tweet.User == user {
			count++
		}
	}
	return count
}

func GetTweetsByUser(user string) []*domain.Tweet {
	return Users[user]
}