package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/derektruong/news-app-gin/database"
	"github.com/gin-gonic/gin"
	// "github.com/gopherjs/gopherjs/js"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AccountFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "account/sign_in_up.html", nil)
}

func LoadEmailToFile(db *sql.DB) {
	type Email struct {
		Email string    `json:"email"`
	}
	// type Data struct {
	// 	Data []Email `json:"data"`
	// }

	data := make([]Email, 0)
	rows, err := db.Query("SELECT a.email FROM NEWS_APP.account a;")

	if err != nil {
		return 
	}

	// var data = Data{ Data: make([]Email, 0)}

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

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" {
		LoadEmailToFile(db)
		c.HTML(http.StatusOK, "account/sign_in_up.html", nil)
		return
	} else {
		// encrypt password
		hash, _ := HashPassword(password)
		fmt.Print(hash)

		_, e := addAccount.Exec(name, email, hash, 3)

		if e != nil {
			// c.HTML(http.StatusOK, "account/sign_in_up.html", gin.H{
			// 	"Name": name,
			// 	"Email": email,
			// 	"Password": password,
			// 	"Show": true,
			// })
			// c.JSON(http.StatusOK, gin.H{
			// 	"message": "Error exist email",
			// })
			
		} else if e == nil {
			
			c.HTML(http.StatusOK, "account/welcome.html", gin.H{
				"Name": name,
			})
		}
	}


	

}
