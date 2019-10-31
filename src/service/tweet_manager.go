package service

import (
	"fmt"
	"github.com/d-bizari/exampleGo/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error{
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	}
	Tweet = tweet
	return nil
}

func GetTweet() *domain.Tweet  {
	return Tweet
}