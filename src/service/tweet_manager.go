package service

import (
	"fmt"
	"github.com/d-bizari/exampleGo/src/domain"
)

var Tweet *domain.Tweet

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
	Tweet = tweet
	return nil
}

func GetTweet() *domain.Tweet {
	return Tweet
}
