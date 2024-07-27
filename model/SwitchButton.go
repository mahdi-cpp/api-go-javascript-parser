package model

import "fmt"

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
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		switchButtons = append(switchButtons, view)
	}
}
