package model_parse

import (
	"encoding/json"
	"fmt"
)

type Button struct {
	ViewBase
	ThemBase
	TextBase
	EventBase
}

var button []Button

func AddButton(jsonString string) {
	var view Button
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error Button unmarshalling JSON:", err)
	} else {
		button = append(button, view)
	}
}
