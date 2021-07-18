package news

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
	// "time"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func (c *Client) FetchSearch(query, page string) (*Result, error) {
	endpoint := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&page=%s&pageSize=%d&apiKey=%s&sortBy=popularity&language=en", url.QueryEscape(query), page, c.PageSize, c.key)
	resp, err := c.http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	var res Result
	json.Unmarshal(body, &res)
	return &res, json.Unmarshal(body, &res)
}

func SearchHandler(cl *Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchQuery := c.Query("q")
		page := c.Query("page")

		if page == "" {
			page = "1"
		}

		results, err := cl.FetchSearch(searchQuery, page)
		// fmt.Println(results)

		if err != nil {
			c.HTML(http.StatusInternalServerError, "general/notfound.html", nil)
			return
		}

		currentPage, err := strconv.Atoi(page)

		if err != nil {
			c.HTML(http.StatusInternalServerError, "test.tmpl", nil)
			return
		}

		resArray := make([][]Article, 0)

		crr := 0

		for {
			mini := int(math.Min(float64(len(results.Articles)), float64(crr + 3)))
			resArray = append(resArray, results.Articles[crr: mini])
			if mini == len(results.Articles) {
				break
			}
			crr += 3
		}

		search := &Search{
			Type: "",
			Path: "",
			Query: searchQuery,
			CurrentPage: currentPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(cl.PageSize))),
			Results: results,
			RowResults: resArray,
		}

		c.HTML(http.StatusOK, "news/news_search.html", search)
	}
	
}