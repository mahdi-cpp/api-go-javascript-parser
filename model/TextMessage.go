package model

import (
	"encoding/json"
	"fmt"
)

type TextMessage struct {
	ViewBase
	TextBase
	BoxColor int `json:"boxColor"`
}

var textMessages []TextMessage

func AddTextMessage(jsonString string) {
	var view TextMessage
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error TextMessage unmarshalling JSON:", err)
	} else {
		textMessages = append(textMessages, view)
	}
}
