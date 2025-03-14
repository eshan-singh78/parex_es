package main

import (
	"fmt"
	"os"
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

	r.POST("/parex_test", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "File not found"})
			return
		}

		err = os.MkdirAll("./tmp", os.ModePerm)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create temp directory"})
			return
		}

		filePath := fmt.Sprintf("./tmp/%s", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}

		c.JSON(200, gin.H{"message": "File uploaded successfully", "file_path": filePath})
	})

	r.Run(":8080")
}
