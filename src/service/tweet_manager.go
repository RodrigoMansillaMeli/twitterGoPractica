package service

import (
	"fmt"
	"github.com/RodrigoMansillaMeli/twitterGoPractica/src/domain"
)

type TweetManager struct {
	Tweets []domain.Tweet
	TweetsByUser map[string][]domain.Tweet
	tweetWriter TweetWriter
}

func NewTweetManager(tweetWriter TweetWriter) *TweetManager {
	tweetManager := new(TweetManager)

	tweetManager.Tweets = make([]domain.Tweet, 0)
	tweetManager.TweetsByUser = make(map[string][]domain.Tweet)
	tweetManager.tweetWriter = tweetWriter

	return tweetManager
}

func (manager *TweetManager) PublishTweet(tweetToPublish domain.Tweet) (int, error){

	if tweetToPublish.GetUser() == "" {
		return tweetToPublish.GetId(), fmt.Errorf("user is required")
	}

	if tweetToPublish.GetText() == "" {
		return tweetToPublish.GetId(), fmt.Errorf("text is required")
	}

	if len(tweetToPublish.GetText()) > 140 {
		return tweetToPublish.GetId(), fmt.Errorf("text exceed 140 characters")
	}

	manager.Tweets = append(manager.Tweets, tweetToPublish)

	tweetToPublish.SetId(len(manager.Tweets))

	manager.TweetsByUser[tweetToPublish.GetUser()] = append(manager.TweetsByUser[tweetToPublish.GetUser()], tweetToPublish)

	manager.tweetWriter.Write(tweetToPublish)

	return tweetToPublish.GetId(), nil
}

func (manager *TweetManager) GetTweet() domain.Tweet {
	return manager.Tweets[len(manager.Tweets)-1]
}

func (manager *TweetManager) GetTweets() []domain.Tweet {
	return manager.Tweets
}

func (manager *TweetManager) GetTweetById(identificador int) domain.Tweet {
	for i:=0; i<len(manager.Tweets); i++ {
		if i+1 == identificador {
			return manager.Tweets[i]
		}
	}
	return nil
}

func (manager *TweetManager) CountTweetsByUser(user string) int {
	counter := 0

	for i:=0; i<len(manager.Tweets); i++ {
		if manager.Tweets[i].GetUser() == user {
			counter++
		}
	}

	return counter
}

func (manager *TweetManager) GetTweetsByUser(user string) map[string][]domain.Tweet {
	return manager.TweetsByUser
}