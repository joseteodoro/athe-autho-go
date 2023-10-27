package routes

import (
	"athe-autho-go/internal/repositories"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterPages(r *gin.Engine) {
	api := r.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			session := sessions.Default(c)
			user := session.Get("user")
			if user != nil {
				c.HTML(http.StatusOK, "index.html", user)
				return
			}

			error := c.DefaultQuery("error", "")
			username := c.DefaultQuery("username", "")
			c.HTML(http.StatusOK, "login.html", struct {
				Error    string
				Username string
			}{Error: error, Username: username})
		})
		api.POST("/login", func(c *gin.Context) {
			handleError := func(err string, username string) {
				c.HTML(http.StatusOK, "login.html", struct {
					Error    string
					Username string
				}{Error: err, Username: username})
			}
			body := struct {
				Username string `binding:"required"`
				Password string `binding:"required"`
				Token    string
			}{}
			if err := c.ShouldBind(&body); err != nil {
				handleError(err.Error(), body.Username)
				return
			}
			user, err := repositories.FindUserByNameAndPassword(body.Username, body.Password)
			if err != nil {
				handleError(err.Error(), body.Username)
				return
			}

			user.Password = "redacted"
			session := sessions.Default(c)
			session.Set("user", user)
			session.Save()

			c.HTML(http.StatusOK, "index.html", user)
		})
		// Add more API routes as needed
	}
}
