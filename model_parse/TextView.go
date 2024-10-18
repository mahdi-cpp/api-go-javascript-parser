package model_parse

import (
	"encoding/json"
	"fmt"
)

type TextView struct {
	ViewBase
	TextBase
}

var textViews []TextView

func AddTextView(jsonString string) {
	var view TextView
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error TextView unmarshalling JSON:", err)
	} else {
		textViews = append(textViews, view)
	}
}
