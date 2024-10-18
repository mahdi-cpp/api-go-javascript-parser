package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Photo struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Orientation int    `json:"orientation"`
	FileType    string `json:"fileType"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type PhotoBase struct {
	Key         int     `json:"key" default:"-1"`
	Name        string  `json:"name"`
	FileType    string  `json:"fileType"`
	Orientation int     `json:"orientation"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Circle      bool    `json:"circle,omitempty"`
	Round       int     `json:"round,omitempty"`
	Crop        bool    `json:"crop,omitempty"`
	AspectRatio float32 `json:"aspectRatio,omitempty"`
	ThumbSize   int     `json:"thumbSize,omitempty"`
	PaintWidth  float32 `json:"paintWidth,omitempty"`
	PaintHeight float32 `json:"paintHeight,omitempty"`
	Dx          float32 `json:"dx,omitempty"`
	Dy          float32 `json:"dy,omitempty"`
}

type IconBase struct {
	Key    int     `json:"key"`
	Name   string  `json:"name"`
	Width  int     `json:"width,omitempty"`
	Height int     `json:"height,omitempty"`
	Dx     float32 `json:"dx,omitempty"`
	Dy     float32 `json:"dy,omitempty"`
	Color  int     `json:"color,omitempty"`
	Alpha  int     `json:"alpha,omitempty"`
}

// IDGenerator is a struct that holds the current ID and a mutex for thread safety
type IDGenerator struct {
	currentID int
	mu        sync.Mutex
}

// NewIDGenerator creates a new IDGenerator instance
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		currentID: 0,
	}
}

// NextID generates the next unique ID
func (g *IDGenerator) NextID() int {
	g.mu.Lock()         // Lock to ensure thread safety
	defer g.mu.Unlock() // Unlock after generating the ID
	g.currentID++       // Increment the current ID
	return g.currentID  // Return the new ID
}

var IdGen = NewIDGenerator()

type PhotoBaseCache struct {
	sync.RWMutex
	Cache map[int]PhotoBase
}

var PhotoBaseMemory = PhotoBaseCache{Cache: make(map[int]PhotoBase)}

func ReadOfFile(folder string, file string) []PhotoBase {

	var photos []PhotoBase

	// Open the file for reading
	f, err := os.Open(folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return photos
}
