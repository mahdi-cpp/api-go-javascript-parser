package model

type ViewBase struct {
	Id     string  `json:"id"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Dx     float32 `json:"dx"`
	Dy     float32 `json:"dy"`
	Round  float32 `json:"round"`
}

type TextBase struct {
	Text      string  `json:"text,omitempty"`
	TextColor int     `json:"textColor,omitempty"`
	TextSize  float32 `json:"textSize,omitempty"`
	Align     string  `json:"align,omitempty"`
}

type ThemBase struct {
	Icon            string `json:"icon,omitempty"`
	Line            bool   `json:"line,omitempty"`
	BackgroundColor int    `json:"backgroundColor,omitempty"`
	ActiveColor     int    `json:"activeColor,omitempty"`
	CircleColor     int    `json:"circleColor,omitempty"`
}
type EventBase struct {
	OnPress string `json:"onPress,omitempty"`
	OnClick string `json:"onClick,omitempty"`
}

// -----------------------------------------

type TextView struct {
	ViewBase
	TextBase
}

type Image struct {
	ViewBase
	Source string `json:"source"`
}

type CircleButton struct {
	ViewBase
	ThemBase
	EventBase
}
type SwitchButton struct {
	ViewBase
	ThemBase
	EventBase
	Checked  bool `json:"checked"`
	Duration int  `json:"duration"`
}
type SliderView struct {
	ViewBase
	ThemBase
	EventBase
	CircleSize float32 `json:"circleSize"`
}

//-----------------------------------------
