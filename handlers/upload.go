package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(c *gin.Context) {
	// Parse the form input for a file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Specify the directory where images will be saved
	saveDir := "./uploads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.Mkdir(saveDir, os.ModePerm)
	}

	// Generate the file path
	filePath := filepath.Join(saveDir, file.Filename)

	// Save the file to the directory
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	// Return the file path or URL
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filePath": filePath,
	})
}
