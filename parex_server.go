package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"parex/internal/lib"

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

	r.POST("/parex_process_v2", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "File not found"})
			return
		}

		offsetStr := c.PostForm("offset")
		levelStr := c.PostForm("level")

		offset, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid offset"})
			return
		}

		level, err := strconv.Atoi(levelStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid level"})
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

		imageFile, err := os.Open(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open file"})
			return
		}
		defer imageFile.Close()

		outputFile := "./tmp/output.log"
		out, err := os.Create(outputFile)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create output log file"})
			return
		}
		defer out.Close()

		startTime := time.Now()

		oldStdout := os.Stdout
		os.Stdout = out
		err = lib.Explore(imageFile, uint64(offset), level)
		os.Stdout = oldStdout

		if err != nil {
			c.JSON(500, gin.H{"error": "Error processing file", "details": err.Error()})
			return
		}

		fileNames, err := lib.ExtractFileNames(outputFile)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to extract filenames", "details": err.Error()})
			return
		}

		err = lib.CleanTmpDirectory("./tmp")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to clean tmp directory", "details": err.Error()})
			return
		}

		timeTaken := time.Since(startTime).Seconds()

		c.JSON(200, gin.H{
			"message":    "Processing completed successfully",
			"filenames":  fileNames,
			"time_taken": fmt.Sprintf("%.2f seconds", timeTaken),
		})
	})

	r.Run(":8080")
}
