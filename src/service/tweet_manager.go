package service

import (
	"fmt"
	"github.com/d-bizari/exampleGo/src/domain"
)

var Tweets []*domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return fmt.Errorf("characters exceeded, only 140 characters are allowed")
	}
	Tweets = append(Tweets, tweet)
	return nil
}

func InitializeService() {
	Tweets = make([]*domain.Tweet, 0)
}

func GetTweets() []*domain.Tweet {
	return Tweets
}

func GetTweet() *domain.Tweet {
	return Tweets[len(Tweets)-1]
}
