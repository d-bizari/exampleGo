package service_test

import (
	"fmt"
	"github.com/d-bizari/exampleGo/src/domain"
	"github.com/d-bizari/exampleGo/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet

	tm := service.NewTweetManager()
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	tm.PublishTweet(tweet)

	// Validation
	publishedTweet := tm.GetTweet()
	assert.Equal(t, publishedTweet.GetUser(), user, "Should be equal")
	assert.Equal(t, publishedTweet.GetText(), text, "Should be equal")
	assert.NotEqual(t, publishedTweet.GetDate(), nil, "Should not be equal")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	var user string
	tm := service.NewTweetManager()
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	tm := service.NewTweetManager()
	var user string = "nana"
	text := ""

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "text is required")
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	var user string = "nana"
	tm := service.NewTweetManager()
	text := "aaaaaaafasdfasdfsdfasfasdfadfjklsfkl;asdjkfaldsjfkalsdfj;asdflkasdlkfaslfasdfksdfaksjfldaklfasjfkasjdkflafsdfasdfasdfadfasfasdfasdfasdfasdfasfdasfasdfafasdfs"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "characters exceeded, only 140 characters are allowed")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet, secondTweet *domain.TextTweet // Fill the tweets with data
	var id1, id2 int64
	user1 := "dani"
	text1 := "hola"
	user2 := "chise"
	text2 := "wolas"

	tweet = domain.NewTextTweet(user1, text1)
	secondTweet = domain.NewTextTweet(user2, text2)

	// Operation
	id1, _ = tm.PublishTweet(tweet)
	id2, _ = tm.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tm.GetTweets()

	assert.Equal(t, len(publishedTweets), 2, "Expected size is 2 but was %d", len(publishedTweets))

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	isValidTweet(t, firstPublishedTweet, id1, user1, text1)
	isValidTweet(t, secondPublishedTweet, id2, user2, text2)
}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int64, user string, text string) {
	assert.Equal(t, tweet.GetText(), text)
	assert.Equal(t, tweet.GetUser(), user)
	assert.Equal(t, tweet.GetId(), id)
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tm := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int64

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tm.PublishTweet(tweet)

	// Validation
	publishedTweet := tm.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)
	// Operation
	count := tm.CountTweetsByUser(user)
	// Validation
	assert.Equal(t, count, 2, "Expected count is 2 but was %d", count)
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	var id1, id2, id3 int64

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	fourthTweet := domain.NewTextTweet(user, "holaaa")
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	// publish the 3 tweets
	id1, _ = tm.PublishTweet(tweet)
	id2, _ = tm.PublishTweet(secondTweet)
	id3, _ = tm.PublishTweet(fourthTweet)
	tm.PublishTweet(thirdTweet)
	// Operation
	tweets := tm.GetTweetsByUser(user)

	for _, tw := range tweets {
		fmt.Println(tw)
	}

	// Validation
	assert.Equal(t, len(tweets), 3)
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
	isValidTweet(t, firstPublishedTweet, id1, user, text)
	isValidTweet(t, secondPublishedTweet, id2, user, secondText)
	isValidTweet(t, fourthTweet, id3, user, "holaaa")
}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	assert.Equal(t, expectedText, text)

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	assert.Equal(t, expectedText, text)
}
