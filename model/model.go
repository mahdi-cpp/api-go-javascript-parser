package model

type ViewBase struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Dx     float32 `json:"dx"`
	Dy     float32 `json:"dy"`
	Round  float32 `json:"round"`
}

type TextBase struct {
	Text      string  `json:"text"`
	TextColor int     `json:"textColor"`
	TextSize  float32 `json:"textSize"`
	Align     string  `json:"align"`
}

type ThemBase struct {
	Icon            string `json:"icon"`
	Line            bool   `json:"line"`
	BackgroundColor int    `json:"backgroundColor"`
	ActiveColor     int    `json:"activeColor"`
	CircleColor     int    `json:"circleColor"`
}
type EventBase struct {
	OnPress string `json:"onPress"`
	OnClick string `json:"onClick"`
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

//-----------------------------------------
