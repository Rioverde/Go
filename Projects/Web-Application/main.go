package main

import (
	"net/http"

	"github.com/Rioverde/Go/Projects/Web-Application/fun"
	"github.com/gin-gonic/gin"
)

func main() {

	//creating default router
	r := gin.Default()

	//parse html files
	r.LoadHTMLGlob("static/*")

	//attach all handler functions to specified pathes
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	})

	r.GET("/login", fun.LoginHandler)
	r.GET("/login/submit", fun.LoginSubmitHandler)

	r.GET("/dashboard", fun.DashboardHandler)

	//run a server on specified port
	r.Run("localhost:2020")
}
