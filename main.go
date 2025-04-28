package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	const viewPath string = "view/"

	router := gin.Default()
	router.Static("/static", viewPath+"static")
	router.LoadHTMLGlob(viewPath + "templates/**/*.html")

	upgrader := websocket.Upgrader{}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about-me.html", gin.H{
			"title": "Carl Eric Doromal - About Me",
		})
	})

	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gin.H{
			"title": "Carl Eric Doromal - Projects",
		})
	})

	router.GET("/international", func(c *gin.Context) {
		c.HTML(http.StatusOK, "international.html", gin.H{
			"title": "Carl Eric Doromal - International",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET("/live-reload", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Websocket Upgrade Error:", err)
			return
		}
		defer conn.Close()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read Error:", err)
				break
			}
		}
	})

	router.Run(":8080")
}
