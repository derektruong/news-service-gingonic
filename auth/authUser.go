package auth

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Error string `json:"error"`
}

func AuthSignInHandler(c *gin.Context) {
	token, _ := c.Cookie("TOKEN_JWT_ID")

	acc, err := VerifyJWT(token)

	if err != nil || token == "" {
		c.JSON(200, gin.H{
			"token": token,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "granted!",
			"uid": acc.ID,
			"name": acc.Name,
			"email": acc.Email,
		})
	}
}

func GetFireBaseToken(c *gin.Context) string {
	token, err := c.Cookie("TOKEN_JWT_ID")

	// fmt.Println(token)
	if err != nil || token == "" {
		// c.JSON(200, gin.H{
		// 	"token":   token,
		// 	"message": err.Error(),
		// })
		return  ""
	}
	return token
}

func AuthLogoutHandler(c *gin.Context) {

	cookie := &http.Cookie{
		Name:     "TOKEN_JWT_ID",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(time.Second),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"message": "logout successfully",
	})
}
