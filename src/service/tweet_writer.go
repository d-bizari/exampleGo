package service

import (
	"github.com/d-bizari/exampleGo/src/domain"
	"os"
)

type TweetWriter interface {
	write(tweet domain.Tweet)
	SetFile(file *os.File)
}

type FileTweeterWriter struct {
	file *os.File
}

func (tw *FileTweeterWriter) write(tweet domain.Tweet) {
	_, _ = tw.file.Write([]byte(tweet.PrintableTweet()))
}

func (ftw *FileTweeterWriter) SetFile(file *os.File) {
	ftw.file = file
}
