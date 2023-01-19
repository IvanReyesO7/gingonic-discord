package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postToDiscordWebHook(params map[string]string) {
	fmt.Println(params)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.POST("/post-to-webhook", func(c *gin.Context) {

		params := make(map[string]string)

		params["webhook_url"] = c.PostForm("webhook_url")
		params["username"] = c.PostForm("username")
		params["content"] = c.PostForm("content")
		params["avatar_url"] = c.PostForm("avatar_url")

		postToDiscordWebHook(params)
	})

	server.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	server.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
