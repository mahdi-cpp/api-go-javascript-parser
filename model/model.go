package model

const (
	VIEW          = 0
	IMAGE         = 1
	TEXT_VIEW     = 2
	CIRCLE_BUTTON = 3
)

type View struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Dx     float32 `json:"dx"`
	Dy     float32 `json:"dy"`
}

// -----------------------------------------

type TextView struct {
	View
	Text      string  `json:"text"`
	TextColor int     `json:"textColor"`
	TextSize  float32 `json:"textSize"`
	Align     string  `json:"align"`
}

type Image struct {
	View
	Source string  `json:"source"`
	Round  float32 `json:"round"`
}

type CircleButton struct {
	View
	Icon    string `json:"icon"`
	OnPress string `json:"onPress"`
	OnClick string `json:"onClick"`
}
type SwitchButton struct {
	View
	Value   int    `json:"value"`
	OnPress string `json:"onPress"`
	OnClick string `json:"onClick"`
}

//-----------------------------------------
