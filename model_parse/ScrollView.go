package model_parse

import (
	"encoding/json"
	"fmt"
)

type ScrollView struct {
	ViewBase
	ThemBase
	TextBase
	EventBase
}

var scrollViews []ScrollView

func AddScrollView(jsonString string) {
	var view ScrollView
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error ScrollView unmarshalling JSON:", err)
	} else {
		scrollViews = append(scrollViews, view)
	}
}
