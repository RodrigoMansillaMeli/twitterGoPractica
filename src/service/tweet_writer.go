package service

import (
	"github.com/RodrigoMansillaMeli/twitterGoPractica/src/domain"
	"os"
)

type TweetWriter interface {
	Write(tweet domain.Tweet)
}

type MemoryTweetWriter struct {
	lastTweet domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}

func (memory *MemoryTweetWriter) Write(tweet domain.Tweet) {
	memory.lastTweet = tweet
}

func (memory *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return memory.lastTweet
}

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {

	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)

	writer.file = file

	return writer
}

func (writer *FileTweetWriter) Write(tweet domain.Tweet) {

	go func() {
		if writer.file != nil {
			byteSlice := []byte(tweet.PrintableTweet() + "\n")
			writer.file.Write(byteSlice)
		}
	}()
}