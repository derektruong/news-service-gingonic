package news

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	http 		*http.Client
	key 		string
	PageSize 	int
}

type Result struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles 	 []Article `json:"articles"`

}

type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Search struct {
	Type string
	Path string
	Query      string
	CurrentPage   int
	TotalPages int
	Results *Result
	RowResults    [][]Article
}

func NewClient(httpClient *http.Client, key string, pageSize int) *Client{
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

func (a *Article) FormatPublishedDate() string {
	year, month, day := a.PublishedAt.Date()

	return fmt.Sprint(day, month, ", ", year)
}

func (s *Search) IsFirstPage() int {
	if s.CurrentPage == 1 {
		return 1
	}
	return s.CurrentPage - 1
}

func (s *Search) IsLastPage() int {
	if s.CurrentPage >= s.TotalPages {
		return s.CurrentPage
	}

	return s.CurrentPage + 1
}