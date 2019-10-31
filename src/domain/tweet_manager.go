package domain

import "time"
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string,text string) *Tweet{
	t := Tweet{user,text,nil}
	tn := time.Now()
	t.Date = &tn
	return &t
}
