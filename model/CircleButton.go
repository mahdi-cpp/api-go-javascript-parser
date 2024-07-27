package model

import (
	"fmt"
)

type CircleButton struct {
	ViewBase
	ThemBase
	EventBase
}

var circleButton []CircleButton

func AddCircleButton(jsonString string) {
	var view CircleButton
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		circleButton = append(circleButton, view)
	}
}
