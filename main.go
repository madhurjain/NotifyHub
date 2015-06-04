package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Notification struct {
	Apn *APN `json:"apn,omitempty"`
	Gcm *GCM `json:"gcm,omitempty"`
	Ws  *WS  `json:"ws,omitempty"`
}

func showHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"title": "NotifyHub"})
}

func sendNotification(c *gin.Context) {
	var notification Notification
	c.BindWith(&notification, binding.JSON)
	err := EnqueueNotification(notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "sent"})
}

func main() {
	// Initialize Notify Hub
	InitHub()

	// Init HTTP Server
	r := gin.Default()
	r.LoadHTMLFiles("home.html")
	r.GET("/", showHome)
	r.POST("/notify", sendNotification)

	r.Run(":8080")
}
