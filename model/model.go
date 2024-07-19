package model

//type Image interface {
//	GetId()
//	GetSource()
//	GetWidth()
//}
//
//type ImageRepository interface {
//	AddImage(image Image)
//	GetAllImage() []Image
//}

type Image struct {
	Source string  `json:"source"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Dx     float32 `json:"dx"`
	Dy     float64 `json:"dy"`
}

type CircleButton struct {
	Icon string `json:"icon"`

	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Dx     float32 `json:"dx"`
	Dy     float64 `json:"dy"`

	OnPress string `json:"onPress"`
	OnClick string `json:"onClick"`
}
