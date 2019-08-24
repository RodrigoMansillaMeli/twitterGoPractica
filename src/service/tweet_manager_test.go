package service_test

import (
	"github.com/RodrigoMansillaMeli/twitterGoPractica/src/domain"
	"github.com/RodrigoMansillaMeli/twitterGoPractica/src/service"
	"strings"
	"testing"
)


func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()
	if publishedTweet.GetUser() != user &&
		publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet

	user := "grupoesfera"
	var text string
	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet

	user := "grupoesfera"
	text := "Este tweet va a tener mas de 140 caracteres para poder hacer una prueba sobre su longitud, ya que si los supera no se podrá publicar por una limitación restrictiva"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text exceed 140 characters" {
		t.Error("Expected error is text exceed 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet *domain.TextTweet

	user := "grupoesfera"
	text := "Primer tweet"
	secondUser := "grupoesfera"
	secondText := "Segundo tweet"

	tweet = domain.NewTextTweet(user,text)
	secondTweet = domain.NewTextTweet(secondUser,secondText)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)
	id2, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		t.Errorf("Invalidad tweet: user or text are not equals")
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id2, secondUser, secondText) {
		t.Errorf("Invalidad tweet: user or text are not equals")
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	if !isValidTweet(t, publishedTweet, id, user, text) {
		t.Errorf("Invalidad tweet: user or text are not equals")
		return
	}
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	// publish the 3 tweets
	id, _ := tweetManager.PublishTweet(tweet)
	id2, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {
		t.Errorf("Expected length is 2 but was %d", len(tweets))
	}

	firstPublishedTweet := tweets[user][0]
	secondPublishedTweet := tweets[user][1]

	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		t.Errorf("Invalidad tweet: user or text are not equals")
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id2, user, secondText) {
		t.Errorf("Invalidad tweet: user or text are not equals")
		return
	}

}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter

	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation

	tweetManager := service.NewTweetManager(tweetWriter)

	tweet := domain.NewTextTweet("grupoesfera", "Texto de tweetWriter")


	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)

	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Error("Tweet was not saved")
		return
	}
	if savedTweet.GetId() != id {
		t.Errorf("Expected tweet id is %d but was %d", savedTweet.GetId(), id)
		return
	}
}

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	// Create and publish a tweet
	tweet := domain.NewTextTweet("grupoesfera", "This is my first tweet")
	secondTweet := domain.NewTextTweet("grupoesfera", "This is my second tweet")
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult //Voy leyendo del canal

	if foundTweet == nil {
		t.Error("Tweet was not founded")
		return
	}
	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Word not found in tweet")
		return
	}

}
func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user string, text string) (bool){
	if tweet.GetUser() != user || tweet.GetText() != text || tweet.GetId() != id {
		return false
	}
	return true
}