package middleware

import (
	"context"
	// "fmt"
	// "fmt"
	"net/http"
	// "strings"

	a "github.com/derektruong/news-app-gin/auth"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations

func AuthMiddleware(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

	// authorizationToken := c.GetHeader("Authorization")

	// fmt.Println(authorizationToken)

	// idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Brearer", "", 1))
	idToken := a.GetFireBaseToken(c)


	
	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Id token not available",
		})
		c.Abort()
		return
	}
	
	// verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid token",
		})
		c.Abort()
		return
	}

	c.Set("UUID", token.UID)
	c.Next()
}
