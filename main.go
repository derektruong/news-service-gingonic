package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/derektruong/news-app-gin/controllers/landing"
	"github.com/derektruong/news-app-gin/controllers/news"
	"github.com/derektruong/news-app-gin/controllers/account"
	"github.com/derektruong/news-app-gin/auth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file!")
	}

	apiKey := os.Getenv("NEWS_API_KEY")

	if apiKey == "" {
		log.Fatal("API key must be set")
	}

	router := gin.Default()
	
	router.Static("/statics/", "./statics")

	router.LoadHTMLGlob("templates/**/*")

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myClient, apiKey, 10)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "general/notfound.html", nil)
	})

	// Handler for API calls
	rAPI := router.Group("/api")
	{
		rAPI.GET("/authuser", auth.AuthSignInHandler)
		rAPI.GET("/authlogout", auth.AuthLogoutHandler)
	}
	
	// account sign in/ up
	router.GET("/signin", account.SignInHandler)
	router.POST("/signin", account.SignInHandler)
	router.GET("/signup", account.SignUpHandler)
	router.POST("/signup", account.SignUpHandler)

	// end account sign in/ up


	router.GET("/", landing.LandingHandler)
	router.GET("/search", news.SearchHandler(newsapi))
	router.GET("/headlines", news.HeadLinesHandler(newsapi))
	router.GET("/stocks", news.StocksHandler(newsapi))
	router.GET("/technology", news.TechHandler(newsapi))
	router.GET("/science", news.ScienceHandler(newsapi))
	router.GET("/sport", news.SportHandler(newsapi))

	router.Run(":3000")


}