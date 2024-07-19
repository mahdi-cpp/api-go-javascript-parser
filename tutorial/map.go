package tutorial

import (
	"fmt"
	"unicode"
)

type Button struct {
	Dx     float64 `json:"dx"`
	Dy     float64 `json:"dy"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Round  float64 `json:"round"`
}

func workingWithMap() {
	var buttons = map[string]Button{}

	buttons["button1"] = Button{Dx: 12, Height: 14, Round: 20}
	buttons["button2"] = Button{Dx: 300, Dy: 40, Width: 300, Height: 70, Round: 20}
	buttons["button3"] = Button{Dx: 50, Dy: 50, Width: 55, Height: 65, Round: 4}

	delete(buttons, "button2")

	for key, element := range buttons {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}
