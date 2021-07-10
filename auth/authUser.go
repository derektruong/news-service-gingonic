package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthSignInHandler(c *gin.Context) {
	cookie := c.Request.Cookies()

	var token string
	for _, val := range cookie {
		if val.Name == "TOKEN_JWT_ID" {
			token = val.Value
		}
	}
	acc, err := VerifyJWT(token)

	if err != nil || token == "" {
		c.JSON(200, gin.H{
			"token": token,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "authenticate successfully!",
			"uid": acc.ID,
			"name": acc.Name,
			"email": acc.Email,
		})
	}
}

func AuthLogoutHandler(c *gin.Context) {

	cookie := &http.Cookie{
		Name: "TOKEN_JWT_ID",
		Value: "",
		Path: "/",
		Expires: time.Now().Add(time.Second),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"message": "logout successfully",
	})
}