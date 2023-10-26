package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type EmptyStruct struct{}

func NewServer() *gin.Engine {
	r := gin.Default()
	// configure session to use cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*.html")
	// page routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./templates/index.html", EmptyStruct{})
	})

	// Implement login route here
	// ...
	return r
}
