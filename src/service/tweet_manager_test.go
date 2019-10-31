package service_test

import (
	"github.com/d-bizari/exampleGo/src/domain"
	"github.com/d-bizari/exampleGo/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	assert.Equal(t, publishedTweet.User, user, "Should be equal")
	assert.Equal(t, publishedTweet.Text, text, "Should be equal")
	assert.NotEqual(t, publishedTweet.Date, nil, "Should not be equal")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string = "nana"
	text := ""

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "text is required")
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string = "nana"
	text := "aaaaaaafasdfasdfsdfasfasdfadfjklsfkl;asdjkfaldsjfkalsdfj;asdflkasdlkfaslfasdfksdfaksjfldaklfasjfkasjdkflafsdfasdfasdfadfasfasdfasdfasdfasdfasfdasfasdfafasdfs"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "characters exceeded, only 140 characters are allowed")
}
