package service_test

import (
	"github.com/d-bizari/exampleGo/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected is", tweet)
	}
}