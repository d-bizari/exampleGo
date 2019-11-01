package domain

import "time"

type Tweet struct {
	User string
	Text string
	Id   int64
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	ptrTweet := &Tweet{user, text, 0, nil}
	tn := time.Now()
	ptrTweet.Date = &tn
	return ptrTweet
}
