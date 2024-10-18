package model_parse

import (
	"encoding/json"
	"fmt"
)

type TextBox struct {
	ViewBase
	TextBase
	BoxColor int `json:"boxColor"`
}

var textBoxes []TextBox

func AddTextBox(jsonString string) {
	var view TextBox
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error TextBox unmarshalling JSON:", err)
	} else {
		textBoxes = append(textBoxes, view)
	}
}
