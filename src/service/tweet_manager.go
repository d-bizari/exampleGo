package service

import "github.com/d-bizari/exampleGo/src/domain"

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet){
	Tweet = tweet
}

func GetTweet() *domain.Tweet  {
	return Tweet
}