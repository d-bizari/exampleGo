package service

import (
	"fmt"
	"github.com/d-bizari/exampleGo/src/domain"
	"os"
)

type TweetManager struct {
	Tweets    []domain.Tweet
	Users     map[string][]domain.Tweet
	IdCounter int64
	Writer    TweetWriter
}

func NewTweetManager() *TweetManager {
	tm := new(TweetManager)
	tm.Tweets = make([]domain.Tweet, 0)
	tm.Users = make(map[string][]domain.Tweet)
	tm.IdCounter = 0
	tm.Writer = new(FileTweeterWriter)
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	tm.Writer.SetFile(file)
	return tm
}

func (tm *TweetManager) PublishTweet(tweet domain.Tweet) (int64, error) {
	if tweet.GetUser() == "" {
		return -1, fmt.Errorf("user is required")
	}
	if tweet.GetText() == "" {
		return -1, fmt.Errorf("text is required")
	}
	if len(tweet.GetText()) > 140 {
		return -1, fmt.Errorf("characters exceeded, only 140 characters are allowed")
	}

	tm.IdCounter++
	tweet.SetId(tm.IdCounter)
	tm.Tweets = append(tm.Tweets, tweet)

	if tm.Users[tweet.GetUser()] == nil {
		tm.Users[tweet.GetUser()] = make([]domain.Tweet, 0)
	}
	tm.Users[tweet.GetUser()] = append(tm.Users[tweet.GetUser()], tweet)
	go tm.Writer.write(tweet)
	return tm.IdCounter, nil
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	return tm.Tweets
}

func (tm *TweetManager) GetTweet() domain.Tweet {
	return tm.Tweets[len(tm.Tweets)-1]
}

func (tm *TweetManager) GetTweetById(id int64) domain.Tweet {
	for _, tweet := range tm.Tweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	count := 0

	for _, tweet := range tm.Tweets {
		if tweet.GetUser() == user {
			count++
		}
	}
	return count
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tm.Users[user]
}
