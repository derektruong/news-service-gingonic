package landing

import (
	// "fmt"
	// "log"
	"net/http"

	// "github.com/derektruong/news-app-gin/auth"
	"github.com/gin-gonic/gin"
)

func QuotesHandler(cl *http.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := FetchQuotes(cl)
		

		if err != nil {
			c.HTML(http.StatusInternalServerError, "general/notfound.html", nil)
			return
		}

		if err != nil {
			c.HTML(http.StatusInternalServerError, "test.tmpl", nil)
			return
		}


		c.HTML(http.StatusOK, "general/index.html", gin.H{
			"Quo1": results.Quotes[0:2],
			"TextCard1": results.Quotes[2].FormatText(),
			"AuthorCard1": results.Quotes[2].Author,
			"TextCard2": results.Quotes[3].FormatText(),
			"AuthorCard2": results.Quotes[3].Author,
			"TextCard3": results.Quotes[4].FormatText(),
			"AuthorCard3": results.Quotes[4].Author,
		})
	}
	
}