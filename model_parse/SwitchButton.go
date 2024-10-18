package model_parse

import (
	"encoding/json"
	"fmt"
)

type SwitchButton struct {
	ViewBase
	ThemBase
	EventBase
	Checked  bool `json:"checked"`
	Duration int  `json:"duration"`
}

var switchButtons []SwitchButton

func AddSwitchButton(jsonString string) {
	var view SwitchButton
	fmt.Println("================")
	fmt.Println(jsonString)
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error SwitchButton unmarshalling JSON:", err)
	} else {
		switchButtons = append(switchButtons, view)
	}
}
