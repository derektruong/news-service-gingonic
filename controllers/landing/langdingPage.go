package landing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QuotesHandler(cl *http.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		results, err := FetchQuotes(cl)
		// fmt.Println(results)

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
			"TextCard1": results.Quotes[2].Text,
			"AuthorCard1": results.Quotes[2].Author,
			"TextCard2": results.Quotes[3].Text,
			"AuthorCard2": results.Quotes[3].Author,
			"TextCard3": results.Quotes[4].Text,
			"AuthorCard3": results.Quotes[4].Author,
		})
	}
	
}