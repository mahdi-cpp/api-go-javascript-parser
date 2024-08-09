package model

import (
	"github.com/gin-gonic/gin"
)

type ViewBase struct {
	Id      string  `json:"id"`
	Width   float32 `json:"width"`
	Height  float32 `json:"height"`
	Dx      float32 `json:"dx"`
	Dy      float32 `json:"dy"`
	Round   float32 `json:"round"`
	Padding float32 `json:"padding,omitempty"`
	Margin  float32 `json:"margin,omitempty"`
	Icon    string  `json:"icon,omitempty"`
}

type TextBase struct {
	Text      string  `json:"text,omitempty"`
	TextColor int     `json:"textColor,omitempty"`
	TextSize  float32 `json:"textSize,omitempty"`
	Align     string  `json:"align,omitempty"`
}

type ThemBase struct {
	Line            bool `json:"line,omitempty"`
	BackgroundColor int  `json:"backgroundColor,omitempty"`
	ActiveColor     int  `json:"activeColor,omitempty"`
	CircleColor     int  `json:"circleColor,omitempty"`
}
type EventBase struct {
	OnPress string `json:"onPress,omitempty"`
	OnClick string `json:"onClick,omitempty"`
}

func Clear() {
	scrollViews = []ScrollView{}
	chartViews = []ChartView{}
	button = []Button{}
	images = []Image{}
	sliderViews = []SliderView{}
	switchButtons = []SwitchButton{}
	textBoxes = []TextBox{}
	textViews = []TextView{}
	textMessages = []TextMessage{}
	textInputs = []TextInput{}
}

func RestAll() map[string]any {
	return gin.H{
		"questionVoices": questionVoices,
		"pdfs":           pdfs,
		"maps":           maps,
		"albums":         albums,
		"gallery":        gallery,
		"movies":         movies,
		"musics":         musics,
		"posts":          posts,
		"users":          users,
		"avatars":        avatars,
		"functions":      RestFunctions(),
		"scrollViews":    scrollViews,
		"textBoxes":      textBoxes,
		"textViews":      textViews,
		"images":         images,
		"buttons":        button,
		"switchButtons":  switchButtons,
		"sliderViews":    sliderViews,
		"chartViews":     chartViews,
		"textMessages":   textMessages,
		"textInputs":     textInputs,
	}
}

func AddView(header string, jsonString string) {

	switch header {
	case "ScrollView":
		AddScrollView(jsonString)
		break
	case "TextView":
		AddTextView(jsonString)
		break
	case "TextBox":
		AddTextBox(jsonString)
		break
	case "Image":
		AddImage(jsonString)
		break
	case "Button":
		AddButton(jsonString)
		break
	case "SwitchButton":
		AddSwitchButton(jsonString)
		break
	case "SliderView":
		AddSliderView(jsonString)
		break
	case "ChartView":
		AddChartView(jsonString)
		break
	case "TextMessage":
		AddTextMessage(jsonString)
		break
	case "TextInput":
		AddTextInput(jsonString)
		break
	}
}
