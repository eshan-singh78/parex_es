package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		loc, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to load timezone"})
			return
		}

		currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05 MST")
		c.JSON(200, gin.H{"message": fmt.Sprintf("Success %s", currentTime)})
	})

	r.Run(":8080")
}
