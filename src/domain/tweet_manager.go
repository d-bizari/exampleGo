package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	PrintableTweet() string
	GetUser() string
	GetText() string
	GetId() int64
	GetDate() *time.Time
	SetId(Id int64)
	String() string
}

type TextTweet struct {
	User string
	Text string
	Id   int64
	Date *time.Time
}

type ImageTweet struct {
	TextTweet
	Url string
}

type QuoteTweet struct {
	TextTweet
	SecondTweet *TextTweet
}

func NewTextTweet(user string, text string) *TextTweet {
	ptrTweet := &TextTweet{user, text, 0, nil}
	tn := time.Now()
	ptrTweet.Date = &tn
	return ptrTweet
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	ptrTweet := &ImageTweet{TextTweet{User: user, Text: text, Id: 0, Date: nil}, url}
	tn := time.Now()
	ptrTweet.Date = &tn
	return ptrTweet
}

func NewQuoteTweet(user string, text string, tweet *TextTweet) *QuoteTweet {
	ptrTweet := &QuoteTweet{TextTweet{User: user, Text: text, Id: 0, Date: nil}, tweet}
	tn := time.Now()
	ptrTweet.Date = &tn
	return ptrTweet
}

func (tw *TextTweet) GetUser() string {
	return tw.User
}

func (tw *TextTweet) GetText() string {
	return tw.Text
}

func (tw *TextTweet) GetId() int64 {
	return tw.Id
}

func (tw *TextTweet) GetDate() *time.Time {
	return tw.Date
}
func (tw *TextTweet) SetId(Id int64) {
	tw.Id = Id
}

func (tw *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%v: %v", tw.User, tw.Text)
}

func (tw *TextTweet) String() string {
	return fmt.Sprintf("@%v: %v", tw.User, tw.Text)
}

func (tw *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%v: %v %v", tw.User, tw.Text, tw.Url)
}
func (tw *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf(`@%v: %v "@%v: %v"`, tw.User, tw.Text, tw.SecondTweet.User, tw.SecondTweet.Text)
}
