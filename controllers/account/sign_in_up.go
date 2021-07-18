package account

import (
	// "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	// "strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/derektruong/news-app-gin/auth"
	"github.com/derektruong/news-app-gin/database"
	"github.com/gin-gonic/gin"

	// "firebase.google.com/go/auth"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TripData(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}


func LoadEmailToFile(db *sql.DB) {
	type Email struct {
		Email string    `json:"email"`
	}

	data := make([]Email, 0)
	rows, err := db.Query("SELECT a.email FROM NEWS_APP.account a;")

	if err != nil {
		return 
	}

	for rows.Next() {
		var email string

		err = rows.Scan(&email)
		

		if err != nil {
			panic(err.Error())
		}
		val := Email{Email: email}
		data = append(data, val)
	}
	fmt.Println(data)
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("./statics/data.json", file, 0644)
}

func SignUpHandler(c *gin.Context) {
	db := database.DBConnect()
	defer db.Close()

	addAccount, err := db.Prepare("INSERT INTO account (name, email, password, idRole) VALUES (?, ?, ?, ?)")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Cannot create account",
		})
	}

	name := TripData(c.PostForm("name"))
	email := TripData(c.PostForm("email"))
	password := TripData(c.PostForm("password"))

	if email == "" {
		LoadEmailToFile(db)
		c.HTML(http.StatusOK, "account/sign_in_up.html", gin.H{
			"Show": true,
		})
		return
	} else {
		// encrypt password
		hash, _ := HashPassword(password)

		addAccount.Exec(name, email, hash, 3)
			
		c.HTML(http.StatusOK, "account/welcome.html", gin.H{
			"Name": name,
		})
	}

}



func SignInHandler(c *gin.Context) {
	db := database.DBConnect()
	defer db.Close()

	email := TripData(c.PostForm("email"))
	password := TripData(c.PostForm("password"))
	isLogin := TripData(c.PostForm("is_login"))

	if email == "" {
		c.HTML(http.StatusOK, "account/sign_in_up.html", nil)
		return
	}

	sqlStatement := "SELECT a.id , a.name, a.password FROM NEWS_APP.account a WHERE email = '"+ email +"'"
	
	var id int
	var name, pass string
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&id, &name, &pass)

	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "Email wrong",
			"text": "This email is not registered, please sign up",
		})

		return
	}
	
	if !CheckPasswordHash(password, pass) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pass wrong",
			"text": "Password is wrong!",
		})
		return
	}

	
	if isLogin == "true" {
		// firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

		// claims := map[string]interface{}{
		// 	"email": email,
		// }

		// token, err := firebaseAuth.CustomTokenWithClaims(context.Background(), strconv.Itoa(id), claims)
		token, err := jwt.Create(id, name, email)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"token": token,
		// 	"message": "Set cookie successfully",
		// })

		expTime := time.Now().AddDate(0, 0, 25)

		cookie := &http.Cookie{
			Name: "TOKEN_JWT_ID",
			Value: token,
			Path: "/",
			Expires: expTime,
			HttpOnly: true,
		}

		http.SetCookie(c.Writer, cookie)



		c.JSON(http.StatusOK, gin.H{
			"message": "Set cookie successfully",
		})
	}
}
