package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

type ImageCache struct {
	sync.RWMutex
	cache map[string][]byte
}

var imageCache = ImageCache{cache: make(map[string][]byte)}

func addToCash(filename string) {

	filepath := "/var/instagram/photos/" + filename // Adjust the path as necessary

	// Read the photo into memory for caching
	imgData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Could not read photo: %s\n", err.Error())
		return
	}

	// Cache the image
	imageCache.Lock()
	imageCache.cache[filename] = imgData
	imageCache.Unlock()
}

func AddDownloadRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/download")
	// Endpoint for downloading a file
	route.GET("/:filename", func(c *gin.Context) {

		filename := c.Param("filename")
		filepath := "/var/instagram/photos/" + filename // Adjust the path as necessary

		fmt.Println("path", filepath)

		// Check if the image is in the cache
		imageCache.RLock()
		imgData, exists := imageCache.cache[filename]
		imageCache.RUnlock()

		// Check if the file exists
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			fmt.Println("File not found")
			c.String(http.StatusNotFound, "File not found")
			return
		}

		if exists {
			// Serve the cached image
			fmt.Println("read of cash")
			c.Data(http.StatusOK, "image/jpeg", imgData) // Adjust MIME type as necessary
		} else {
			// Serve the file for download
			fmt.Println("read of disk drive")
			c.File(filepath)
			addToCash(filename)
		}

	})
}
