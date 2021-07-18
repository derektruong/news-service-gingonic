package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/derektruong/news-app-gin/auth"
	"github.com/derektruong/news-app-gin/config"
	"github.com/derektruong/news-app-gin/controllers/account"
	"github.com/derektruong/news-app-gin/controllers/landing"
	"github.com/derektruong/news-app-gin/controllers/news"
	// "github.com/derektruong/news-app-gin/middleware"

	// "github.com/derektruong/news-app-gin/middleware"

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

	
	

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "general/notfound.html", nil)
	})

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	// set db to gin context with a middleware to all incoming request
	router.Use(func (c *gin.Context){
		c.Set("firebaseAuth", firebaseAuth)
	})

	// Handler for API calls
	rAPI := router.Group("/api") 
	{
		// rAPI.GET("/auth", middleware.AuthMiddleware)
		rAPI.GET("/authuser", auth.AuthSignInHandler)
		rAPI.GET("/authlogout", auth.AuthLogoutHandler)
	}

	
	// account sign in/ up
	router.GET("/signin", account.SignInHandler)
	router.POST("/signin", account.SignInHandler)

	router.GET("/signup", account.SignUpHandler)
	router.POST("/signup", account.SignUpHandler)
	// end account sign in/ up

	// Handler for landing Page use API
	router.GET("/", landing.QuotesHandler(myClient))

	// authorized := router.Group("/")
	// // per group middleware! in this case we use the custom created
	// // AuthRequired() middleware just in the "authorized" group.
	// authorized.Use(middleware.AuthMiddleware)
	// {
		
	// }
	

	// Handler for News Pages use API
	newsapi := news.NewClient(myClient, apiKey, 12)

	router.GET("/search", news.SearchHandler(newsapi))
	router.GET("/headlines", news.HeadLinesHandler(newsapi))
	router.GET("/stocks", news.StocksHandler(newsapi))
	router.GET("/technology", news.TechHandler(newsapi))
	router.GET("/science", news.ScienceHandler(newsapi))
	router.GET("/sport", news.SportHandler(newsapi))

	router.Run(":3000")


}