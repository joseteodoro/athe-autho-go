package server

import (
	"athe-autho-go/api/routes"
	"athe-autho-go/internal/configs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type EmptyStruct struct{}

func NewServer() *gin.Engine {
	config := configs.NewSecurityConfig()
	r := gin.Default()
	// configure session to use cookies
	store := cookie.NewStore([]byte(config.CookieSecret))
	r.Use(sessions.Sessions("hypersession", store))

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	routes.RegisterAPIs(r)
	routes.RegisterPages(r)
	return r
}
