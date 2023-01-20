package main

import (
	"fmt"
	"net/http"

	Discord "gingonic-discord/discord"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.Static("/assets", "./assets")
	server.LoadHTMLGlob("templates/index.html")
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("session", store))

	server.POST("/post-to-webhook", func(c *gin.Context) {
		session := sessions.Default(c)

		webhook_url := c.PostForm("webhook_url")
		params := make(map[string]string)

		params["username"] = c.PostForm("username")
		params["content"] = c.PostForm("content")
		params["avatar_url"] = c.PostForm("avatar_url")

		response := Discord.PostToDiscordWebHook(webhook_url, params)

		if response != "" {
			fmt.Println(response)
			session.AddFlash(response)
			session.Save()
			c.JSON(http.StatusInternalServerError, "")
		} else {
			c.JSON(http.StatusOK, "")
		}
	})

	server.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"alert": session.Flashes(),
		})
	})

	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
