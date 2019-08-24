package domain

import "time"

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(int)
	PrintableTweet() string
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id int
}

type ImageTweet struct {
	TextTweet
	URL string
}

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func NewTextTweet(user, text string) *TextTweet {
	date := time.Now()

	tweet := TextTweet{
		User : user,
		Text: text,
		Date: &date,
	}

	return &tweet
}


func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) PrintableTweet() string {
	return "@"+tweet.User+": "+tweet.Text
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func NewImageTweet(user, text, url string) *ImageTweet {
	date := time.Now()
	imageTweet := ImageTweet{
		TextTweet: TextTweet{
			User : user,
			Text: text,
			Date: &date,
		},
		URL: url,
	}

	return &imageTweet
}

func (tweet *ImageTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *ImageTweet) PrintableTweet() string {
	return "@"+tweet.User+": "+tweet.Text+" "+tweet.URL
}

func (tweet *ImageTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *ImageTweet) GetId() int {
	return tweet.Id
}

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet {
	date := time.Now()

	quoteTweet := QuoteTweet{
		TextTweet: TextTweet{
			User : user,
			Text: text,
			Date: &date,
		},
		QuotedTweet: quotedTweet,
	}

	return &quoteTweet
}

func (tweet *QuoteTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return "@"+tweet.User+": "+tweet.Text+" \"@"+tweet.QuotedTweet.GetUser()+": "+tweet.QuotedTweet.GetText()+"\""
}

func (tweet *QuoteTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *QuoteTweet) GetId() int {
	return tweet.Id
}

