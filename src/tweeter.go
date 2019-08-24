package main

import (
	"github.com/RodrigoMansillaMeli/twitterGoPractica/src/domain"
	"github.com/RodrigoMansillaMeli/twitterGoPractica/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/abiosoft/ishell"

)

var tweetManager *service.TweetManager

// JSONTweet structure to receive tweets
type JSONTweet struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	URL       string `json:"url"`
	IDMencion string `json:"id"`
}

func main() {

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager = service.NewTweetManager(tweetWriter)

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTextTweet(user, text)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes a tweet with an image",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			c.Print("Type the url of your image: ")

			url := c.ReadLine()

			tweet := domain.NewImageTweet(user, text, url)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishQuoteTweet",
		Help: "Publishes a tweet with a quote",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			c.Print("Type the id of the tweet you want to quote: ")

			id, _ := strconv.Atoi(c.ReadLine())

			quoteTweet := tweetManager.GetTweetById(id)

			tweet := domain.NewQuoteTweet(user, text, quoteTweet)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows the last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet with the provided id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the id: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet := tweetManager.GetTweetById(id)

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts the tweets published by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the user: ")

			user := c.ReadLine()

			count := tweetManager.CountTweetsByUser(user)

			c.Println(count)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows the tweets published by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the user: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			c.Println(tweets)

			return
		},
	})

	go levantarAPI()

	shell.Run()

}

func levantarAPI() {
	router := gin.Default()

	router.GET("/tweet", getLastTweet)
	router.GET("/tweets", getAllTweets)
	router.GET("/tweet/:id", getTweetById)
	router.GET("/tweets/:user", getTweetsByUser)

	router.POST("/tweet", publishTweet)

	router.Run()
}

func getLastTweet(c *gin.Context) {
	c.JSON(http.StatusOK, tweetManager.GetTweet())
}

func getAllTweets(c *gin.Context) {
	c.JSON(http.StatusOK, tweetManager.GetTweets())
}

func getTweetById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	c.JSON(http.StatusOK, tweetManager.GetTweetById(id))
}

func getTweetsByUser(c *gin.Context) {
	user := c.Param("user")
	c.JSON(http.StatusOK, tweetManager.GetTweetsByUser(user))
}

func publishTweet(c *gin.Context) {
	var tweetToPublish JSONTweet
	if err := c.ShouldBindJSON(&tweetToPublish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet := getTweetOfItsType(tweetToPublish)
	tweetManager.PublishTweet(tweet)
	c.JSON(http.StatusOK, tweetManager.GetTweet())
}

func getTweetOfItsType(jsonTweet JSONTweet) domain.Tweet {
	switch {
	case jsonTweet.URL != "":
		return domain.NewImageTweet(jsonTweet.User, jsonTweet.Text, jsonTweet.URL)
	case jsonTweet.IDMencion != "":
		id, _ := strconv.Atoi(jsonTweet.IDMencion)
		return domain.NewQuoteTweet(jsonTweet.User, jsonTweet.Text, tweetManager.GetTweetById(id))
	default:
		return domain.NewTextTweet(jsonTweet.User, jsonTweet.Text)
	}
}