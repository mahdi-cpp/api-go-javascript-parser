package model

import "fmt"

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
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		textBoxes = append(textBoxes, view)
	}
}
