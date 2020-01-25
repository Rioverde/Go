package fun

import (
	"net/http"

	database "github.com/Rioverde/Go/Projects/Web-Application/crud"
	"github.com/gin-gonic/gin"
)

var err error

// LoginHandler handles upcoming requests on /login path
func LoginHandler(c *gin.Context) {
	//if cookie is set then redirect to dashboard
	val, err := c.Cookie("email")
	if err == nil && val != "" {
		c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
		return
	}

	//write to the response stream login.html contents
	c.HTML(http.StatusOK, "extra.html", nil)
}

// LoginSubmitHandler processes form data from login page
func LoginSubmitHandler(c *gin.Context) {
	//searching and getting email and password from http query
	email := c.Query("email")
	password := c.Query("password")

	_, err = database.CheckUser(email, password)
	if err == nil {
		//set cookie email=value for all pages of server
		c.SetCookie("email", email, 86400, "", "127.0.0.1", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

//DashboardHandler handles upcoming requests on /dashboard path
func DashboardHandler(c *gin.Context) {
	//check if cookie is set, if no then redirect to /login path
	val, err := c.Cookie("email")
	if err != nil || val == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	c.HTML(http.StatusOK, "dash.html", val)
}
