package model

import "fmt"

type TextView struct {
	ViewBase
	TextBase
}

var textViews []TextView

func AddTextView(jsonString string) {
	var view TextView
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		textViews = append(textViews, view)
	}
}
