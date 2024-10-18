package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

//var root = "/var/cloud/bb/"
//var rootThumb = "/var/cloud/bb/thumbnail/"

var folders = []string{
	"/var/cloud/bb/",
	"/var/cloud/fa/",
	"/var/cloud/wallpaper/",

	"/var/instagram/id/ali/",
	"/var/instagram/id/cut/",
	"/var/instagram/id/girl/",
	"/var/instagram/id/face/",
	"/var/instagram/id/trip/",
	"/var/instagram/id/go/",

	"/var/instagram/chat/movie/movie/",
	"/var/instagram/chat/pdf/",
	"/var/instagram/chat/electronic/",
	"/var/instagram/chat/map/",
	"/var/instagram/chat/voice/",

	"/var/instagram/call2/",
}
var thumbFolders = []string{
	"/var/cloud/bb/thumbnail/",
	"/var/cloud/fa/thumbnail/",
	"/var/instagram/id/ali/thumbnail/",
	"/var/instagram/id/cut/thumbnail/",
	"/var/instagram/id/girl/thumbnail/",
	"/var/instagram/id/face/thumbnail/",
	"/var/instagram/id/trip/thumbnail/",
	"/var/instagram/id/go/thumbnail/",

	"/var/instagram/chat/movie/movie/thumbnail/",
	"/var/instagram/chat/pdf/thumbnail/",
	"/var/instagram/chat/electronic/thumbnail/",
	"/var/instagram/chat/map/thumbnail/",
	"/var/instagram/chat/voice/thumbnail/",

	"/var/instagram/call2/thumbnail/",
}

var iconFolder = "/var/cloud/icons/"

type ImageCache struct {
	sync.RWMutex
	cache map[string][]byte
}

var imageCache = ImageCache{cache: make(map[string][]byte)}

type IconCache struct {
	sync.RWMutex
	cache map[string]image.Image
}

var iconCache = IconCache{cache: make(map[string]image.Image)}

func searchFile(folders []string, filename string) (string, error) {
	for _, folder := range folders {
		// Construct the full path to the file
		fullPath := filepath.Join(folder, filename)

		// Check if the file exists
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil // File found
		} else if os.IsNotExist(err) {
			// File does not exist in this directory
			continue
		} else {
			// Other error (e.g., permission issues)
			return "", err
		}
	}
	return "", fmt.Errorf("file %s not found in any of the specified folders", filename)
}

// LoadImage loads an image from a file.
func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func addIconCash(iconName string) {
	icon, err := LoadImage("/var/cloud/icons/" + iconName) // Change path accordingly
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	iconCache.Lock()
	iconCache.cache[iconName] = icon
	iconCache.Unlock()
}

func addToCash(filepath string, filename string) {

	originalImage, err := LoadImage(filepath)
	if err != nil {
		fmt.Println("addToCash Error loading image:", err)
		return
	}

	imgBytes, err := ConvertImageToBytes(originalImage, "jpg") // Change to "png" for PNG format
	if err != nil {
		fmt.Println("Error ConvertImageToBytes: ", err)
		return
	}

	imageCache.Lock()
	imageCache.cache[filename] = imgBytes
	imageCache.Unlock()
}

// ConvertImageToBytes converts an image.Image to a byte slice.
func ConvertImageToBytes(img image.Image, format string) ([]byte, error) {
	var buf bytes.Buffer
	var err error

	// Encode the image based on the specified format
	switch format {
	case "jpg":
		err = jpeg.Encode(&buf, img, nil)
	case "png":
		err = png.Encode(&buf, img)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func AddDownloadRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/download")
	apiOriginalDownload(route)
	apiDownloadThumb(route)
	apiIcon(route)
}

func apiOriginalDownload(route *gin.RouterGroup) {

	route.GET("/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filepath, err := searchFile(folders, filename)
		if err != nil {
			return
		}
		c.File(filepath)
	})
}

func apiDownloadThumb(route *gin.RouterGroup) {

	route.GET("/thumbnail/:filename", func(c *gin.Context) {

		filename := c.Param("filename")

		imageCache.RLock()
		imgData, exists := imageCache.cache[filename]
		imageCache.RUnlock()

		if exists {
			fmt.Println("RAM")
			c.Data(http.StatusOK, "image/jpeg", imgData) // Adjust MIME type as necessary
		} else {

			filepath, err := searchFile(thumbFolders, filename)
			if err != nil {
				return
			}

			// Serve the file for download
			fmt.Println("SSD")
			c.File(filepath)
			addToCash(filepath, filename)
		}
	})
}

func apiIcon(route *gin.RouterGroup) {

	route.GET("/icons/:filename", func(c *gin.Context) {

		filename := c.Param("filename")

		iconCache.RLock()
		imgData, exists := iconCache.cache[filename]
		iconCache.RUnlock()

		if exists {
			imgBytes, err := ConvertImageToBytes(imgData, "png") // Change to "png" for PNG format
			if err != nil {
				fmt.Println("Error converting image to bytes:", err)
				return
			}
			c.Data(http.StatusOK, "image/png", imgBytes) // Adjust MIME type as necessary
		}
	})
}

func ReadIcons() {
	// Specify the directory you want to read
	dir := "/var/cloud/icons" // Change this to your target directory

	// Read the directory entries
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	// Iterate over the entries
	for _, entry := range entries {
		if !entry.IsDir() { // Check if it is not a directory

			if strings.Contains(entry.Name(), ".png") {
				addIconCash(entry.Name())
				//fmt.Printf("Reading file: %s\n", entry.Name())
			}
		}
	}

	fmt.Println(len(iconCache.cache))
}
