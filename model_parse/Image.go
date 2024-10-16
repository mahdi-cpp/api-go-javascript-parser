package model_parse

import (
	"encoding/json"
	"fmt"
)

type Image struct {
	ViewBase
	Source string `json:"source"`
}

var images []Image

func AddImage(jsonString string) {
	var view Image
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error AddImage unmarshalling JSON:", err)
	} else {
		images = append(images, view)
	}
}
