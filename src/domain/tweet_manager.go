package domain

import "time"
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string,text string) *Tweet{
	ptrTweet := &Tweet{user,text,nil}
	/*
	ptrTweet := new(Tweet)
	ptrTweet.User = user
	ptrTweet.Text = text
	*/
	tn := time.Now()
	ptrTweet.Date = &tn
	return ptrTweet
}
