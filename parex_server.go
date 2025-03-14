package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	r.POST("/parex_process", func(c *gin.Context) {
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

		err = lib.Explore(imageFile, uint64(offset), level)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error processing file", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Processing completed successfully"})
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

		// Redirect stdout to capture output
		oldStdout := os.Stdout
		os.Stdout = out
		err = lib.Explore(imageFile, uint64(offset), level)
		os.Stdout = oldStdout

		if err != nil {
			c.JSON(500, gin.H{"error": "Error processing file", "details": err.Error()})
			return
		}

		// Read and extract filenames from output
		fileNames, err := extractFileNames(outputFile)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to extract filenames", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message":   "Processing completed successfully",
			"filenames": fileNames,
		})
	})

	r.Run(":8080")
}

// extractFileNames extracts filenames from the process output
func extractFileNames(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileNames []string
	scanner := bufio.NewScanner(file)

	// Regex to capture filenames at the end of each line
	re := regexp.MustCompile(`\S+\.\w+$`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindString(line)
		if matches != "" {
			fileNames = append(fileNames, matches)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileNames, nil
}
