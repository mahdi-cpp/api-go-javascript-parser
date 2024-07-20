package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"strings"
)

func parseViewBase(viewParse ViewParse) model.ViewBase {
	var dto model.ViewBase
	var index = 0

	for _, attr := range viewParse.Attributes {
		switch attr {
		case "dx":
			dto.Dx = utils.GetFloat(viewParse.Values[index])
			break
		case "dy":
			dto.Dy = utils.GetFloat(viewParse.Values[index])
			break
		case "width":
			dto.Width = utils.GetFloat(viewParse.Values[index])
			break
		case "height":
			dto.Height = utils.GetFloat(viewParse.Values[index])
			break
		case "round":
			dto.Round = utils.GetFloat(viewParse.Values[index])
			break
		}
		index++
	}

	return dto
}

func parseTextBase(viewParse ViewParse) model.TextBase {
	var dto model.TextBase
	var index = 0

	for _, attr := range viewParse.Attributes {

		switch attr {
		case "text":
			var t = viewParse.Values[index]
			t = strings.Replace(t, "'", "", 2)
			dto.Text = t
			break
		case "textColor":
			dto.TextColor = utils.GetColor(viewParse.Values[index])
			break
		case "textSize":
			dto.TextSize = utils.GetFloat(viewParse.Values[index])
			break
		case "align":
			var t = viewParse.Values[index]
			t = strings.Replace(t, "'", "", 2)
			dto.Align = t
			break
		}
		index++
	}

	return dto
}

func ParserThem(viewParse ViewParse) model.ThemBase {
	var dto model.ThemBase
	var index = 0

	for _, attr := range viewParse.Attributes {
		switch attr {
		case "line":
			dto.Line = ParseBoolean(viewParse.Values[index])
			break
		case "backgroundColor":
			dto.BackgroundColor = utils.GetColor(viewParse.Values[index])
			break
		case "activeColor":
			dto.ActiveColor = utils.GetColor(viewParse.Values[index])
			break
		case "circleColor":
			dto.CircleColor = utils.GetColor(viewParse.Values[index])
			break
		case "icon":
			var t = viewParse.Values[index]
			t = strings.Replace(t, "'", "", 2)
			dto.Icon = t
			break
		}
		index++
	}
	return dto
}

func ParserEventBase(viewParse ViewParse) model.EventBase {
	var dto model.EventBase
	var index = 0

	for _, attr := range viewParse.Attributes {
		switch attr {
		case "onPress":
			dto.OnPress = viewParse.Values[index]
			break
		case "onClick":
			dto.OnClick = viewParse.Values[index]
			break
		}
		index++
	}
	return dto
}

//-----------------------------------------

func TextParser(viewParse ViewParse) {
	var dto model.TextView
	dto.ViewBase = parseViewBase(viewParse)
	dto.TextBase = parseTextBase(viewParse)
	textViews = append(textViews, dto)
}

func ImageParser(viewParse ViewParse) {
	var dto model.Image
	dto.ViewBase = parseViewBase(viewParse)
	var index = 0
	for _, attr := range viewParse.Attributes {
		switch attr {
		case "source":
			var t = viewParse.Values[index]
			t = strings.Replace(t, "'", "", 2)
			dto.Source = t
			break
		}
		index++
	}

	fmt.Println(dto)
	images = append(images, dto)
}

func SwitchButtonParser(viewParse ViewParse) {
	var dto model.SwitchButton
	dto.ViewBase = parseViewBase(viewParse)
	dto.ThemBase = ParserThem(viewParse)
	dto.EventBase = ParserEventBase(viewParse)

	var index = 0

	for _, attr := range viewParse.Attributes {
		switch attr {
		case "checked":
			dto.Checked = ParseBoolean(viewParse.Values[index])
			break
		case "duration":
			dto.Duration = ParseInteger(viewParse.Values[index])
			break
		}
		index++
	}
	switchButtons = append(switchButtons, dto)
}

func CircleButtonParser(viewParse ViewParse) {

	var dto model.CircleButton
	dto.ViewBase = parseViewBase(viewParse)
	dto.ThemBase = ParserThem(viewParse)
	dto.EventBase = ParserEventBase(viewParse)

	circleButtons = append(circleButtons, dto)
}
