package landing

import (
	// "math"
	"net/http"
	// "net/url"
	// "strconv"
	// "time"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"

	"github.com/gin-gonic/gin"
)

func LandingHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "general/index.html", nil)
	
}