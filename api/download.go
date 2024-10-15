package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

//var root = "/var/cloud/bb/"
//var rootThumb = "/var/cloud/bb/thumbnail/"

var folders = []string{
	"/var/cloud/bb/",
	"/var/cloud/fa/",
	"/var/instagram/id/ali/",
	"/var/instagram/id/cut/",
	"/var/instagram/id/girl/",
}
var thumbFolders = []string{
	"/var/cloud/bb/thumbnail/",
	"/var/cloud/fa/thumbnail/",
	"/var/instagram/id/ali/thumbnail/",
	"/var/instagram/id/cut/thumbnail/",
	"/var/instagram/id/girl/thumbnail/",
}

type ImageCache struct {
	sync.RWMutex
	cache map[string]image.Image
}

var imageCache = ImageCache{cache: make(map[string]image.Image)}

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

func addToCash(filename string) {

	//filepath := rootThumb + filename // Adjust the path as necessary

	// Load the original image
	originalImage, err := LoadImage(filename) // Change path accordingly
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Define the rectangle to crop (x, y, width, height)
	//cropRect := image.Rect(0, 0, 540, 540) // Change coordinates as needed

	// Crop the image
	//croppedImage := CropImage(originalImage, cropRect)

	// Define the rectangle to crop (x, y, width, height)
	//cropRect := image.Rect(0, 0, 270, 270) // Change coordinates as needed

	// Crop the image
	//croppedImage := imaging.Crop(originalImage, cropRect)

	// Save the cropped image to a file
	//err = SaveImage("/var/cloud/crop/"+filename, croppedImage) // Change path accordingly
	//if err != nil {
	//	fmt.Println("Error saving image:", err)
	//	return
	//}

	//// Read the photo into memory for caching
	//imgData, err := os.ReadFile(filepath)
	//if err != nil {
	//	fmt.Printf("Could not read photo: %s\n", err.Error())
	//	return
	//}

	// Cache the image
	imageCache.Lock()
	imageCache.cache[filename] = originalImage
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
	apiDownload(route)
	apiDownloadThumb(route)
}

func apiDownload(route *gin.RouterGroup) {

	route.GET("/:filename", func(c *gin.Context) {

		filename := c.Param("filename")
		//filepath := root + filename // Adjust the path as necessary

		//fmt.Println("original path", filepath)

		// Check if the image is in the cache
		imageCache.RLock()
		imgData, exists := imageCache.cache[filename]
		imageCache.RUnlock()

		if exists {
			imgBytes, err := ConvertImageToBytes(imgData, "jpg") // Change to "png" for PNG format
			if err != nil {
				fmt.Println("Error converting image to bytes:", err)
				return
			}
			// Serve the cached image
			fmt.Println("RAM")
			c.Data(http.StatusOK, "image/jpeg", imgBytes) // Adjust MIME type as necessary
		} else {

			//// Check if the file exists
			//if _, err := os.Stat(filepath); os.IsNotExist(err) {
			//	fmt.Println("File not found")
			//	c.String(http.StatusNotFound, "File not found")
			//	return
			//}

			filepath, err := searchFile(folders, filename)
			if err != nil {
				return
			}

			// Serve the file for download
			fmt.Println("SSD")
			c.File(filepath)
			//addLargeToCash(filename)
		}

	})
}

func apiDownloadThumb(route *gin.RouterGroup) {

	route.GET("/thumbnail/:filename", func(c *gin.Context) {

		filename := c.Param("filename")
		//filepath := rootThumb + filename // Adjust the path as necessary

		//fmt.Println("thumbnail path", filepath)

		// Check if the image is in the cache
		imageCache.RLock()
		imgData, exists := imageCache.cache[filename]
		imageCache.RUnlock()

		if exists {
			imgBytes, err := ConvertImageToBytes(imgData, "jpg") // Change to "png" for PNG format
			if err != nil {
				fmt.Println("Error converting image to bytes:", err)
				return
			}
			// Serve the cached image
			fmt.Println("RAM")
			c.Data(http.StatusOK, "image/jpeg", imgBytes) // Adjust MIME type as necessary
		} else {

			//// Check if the file exists
			//if _, err := os.Stat(filepath); os.IsNotExist(err) {
			//	fmt.Println("File not found")
			//	c.String(http.StatusNotFound, "File not found")
			//	return
			//}

			filepath, err := searchFile(thumbFolders, filename)
			if err != nil {
				return
			}

			// Serve the file for download
			fmt.Println("SSD")
			c.File(filepath)
			addToCash(filename)
		}
	})
}

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

//func addLargeToCash(filename string) {
//
//	filepath := root + filename // Adjust the path as necessary
//
//	// Load the original image
//	originalImage, err := LoadImage(filepath) // Change path accordingly
//	if err != nil {
//		fmt.Println("Error loading image:", err)
//		return
//	}
//
//	imageCache.Lock()
//	imageCache.cache[filename] = originalImage
//	imageCache.Unlock()
//}
