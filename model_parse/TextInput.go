package model_parse

import (
	"encoding/json"
	"fmt"
)

type TextInput struct {
	ViewBase
	TextBase
}

var textInputs []TextInput

func AddTextInput(jsonString string) {
	var view TextInput
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error TextInput unmarshalling JSON:", err)
	} else {
		textInputs = append(textInputs, view)
	}
}
