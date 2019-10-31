package service_test

import (
	"github.com/d-bizari/exampleGo/src/domain"
	"github.com/d-bizari/exampleGo/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T){
	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	assert.Equal(t,publishedTweet.User,user,"Should be equal")
	assert.Equal(t,publishedTweet.Text,text,"Should be equal")
	assert.NotEqual(t,publishedTweet.Date,nil,"Should not be equal")
}