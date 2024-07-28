package model

import (
	"encoding/json"
	"fmt"
)

type SliderView struct {
	ViewBase
	ThemBase
	EventBase
	CircleSize float32 `json:"circleSize"`
}

var sliderViews []SliderView

func AddSliderView(jsonString string) {
	var view SliderView
	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		sliderViews = append(sliderViews, view)
	}
}
