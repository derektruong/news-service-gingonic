package news

import (
	"math"
	"net/http"
	"strconv"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func (c *Client) FetchTech(page string) (*Result, error) {
	endpoint := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&category=technology&page=%s&pageSize=%d&apiKey=%s&sortBy=popularity&language=en", page, c.PageSize, c.key)
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
	fmt.Println(res)
	return &res, json.Unmarshal(body, &res)
}

func TechHandler(cl *Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Query("page")

		if page == "" {
			page = "1"
		}

		results, err := cl.FetchTech(page)
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

		search := &Search{
			Type: "Technologies",
			Path: "technology",
			Query: "",
			CurrentPage: currentPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(cl.PageSize))),
			Results: results,
		}
		c.HTML(http.StatusOK, "news/category.html", search)
	}
	
}