package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
)

const DASHBOARD_UID = "dashboard"

type Notification struct {
	Apn *APN `json:"apn,omitempty"`
	Gcm *GCM `json:"gcm,omitempty"`
}

func showHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"uid": DASHBOARD_UID})
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

func userDashboard(c *gin.Context) {
	uid := c.Param("uid")
	c.HTML(http.StatusOK, "user.html", gin.H{"uid": uid})
}

func wsHandler(c *gin.Context) {
	uid := c.Param("uid")
	ServeWs(c.Writer, c.Request, uid)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// Initialize Notify Hub
	InitHub()

	// Initialize Web Sockets
	InitWebSockets()

	// Init HTTP Server
	r := gin.Default()

	r.LoadHTMLFiles("templates/home.html", "templates/user.html")
	r.Static("/assets", "./assets")

	r.GET("/", showHome)
	r.POST("/notify", sendNotification)
	r.GET("/user/:uid", userDashboard)
	r.GET("/ws/:uid", wsHandler)
	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	r.Run(httpHost + ":" + httpPort)
}
