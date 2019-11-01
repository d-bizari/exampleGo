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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "characters exceeded, only 140 characters are allowed")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data
	var id1, id2 int64
	user1 := "dani"
	text1 := "hola"
	user2 := "chise"
	text2 := "wolas"

	tweet = domain.NewTweet(user1, text1)
	secondTweet = domain.NewTweet(user2, text2)

	// Operation
	id1, _ = service.PublishTweet(tweet)
	id2, _ = service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	assert.Equal(t, len(publishedTweets), 2, "Expected size is 2 but was %d", len(publishedTweets))

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	isValidTweet(t, firstPublishedTweet, id1, user1, text1)
	isValidTweet(t, secondPublishedTweet, id2, user2, text2)
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int64, user string, text string) {
	assert.Equal(t, tweet.Text, text)
	assert.Equal(t, tweet.User, user)
	assert.Equal(t, tweet.Id, id)
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int64

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	assert.Equal(t, count, 2, "Expected count is 2 but was %d", count)
}
