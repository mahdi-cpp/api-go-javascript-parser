package repository

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"regexp"
	"strconv"
	"strings"
)

type View2 struct {
	Header     string
	Properties []Property
	Lines      []string
}

func StartParseViews() {
	model.Clear()
	FindViews("web/index.js")
}

func FindViews(fileName string) {
	fmt.Println("======================================")
	fmt.Println("ParseViews")
	jsonString, err := readJsonFile(fileName)
	if err != nil {
		return
	}
	jsonString = RemoveMultiLineComments(jsonString)
	jsonString = RemoveLineComments(jsonString)

	var search = "<View>"
	_, after, found := strings.Cut(jsonString, search)

	if found {
		fmt.Println(after)
	}
	if found {
		var insideSearch = "</View>"
		before, _, hasEndView := strings.Cut(after, insideSearch)
		if hasEndView {
			views := ParseViews(before)
			views = ParseProperties(views)
			EndParseViews(views)

		} else {
			fmt.Println("Not Found: " + "'" + search + "'")
		}
	} else {
		fmt.Println("Not Found: " + "'" + search + "'")
	}
}

func ParseViews(jsonViews string) []View2 {

	var views []View2

	re := regexp.MustCompile(`<\s*(\w+)\s*`)
	var headers []string

	// Find all sub matches in the input string
	matches := re.FindAllStringSubmatch(jsonViews, -1)
	for _, match := range matches {
		if len(match) > 1 {
			headers = append(headers, match[1])
		}
	}

	result := splitStringWithHeaders(jsonViews, headers)
	for _, section := range result {
		lines := strings.Split(section, "\n")
		var view View2
		for _, line := range lines {

			if strings.HasPrefix(line, "<") {
				view.Header = line
			} else if strings.HasPrefix(line, "/>") {
				break
			} else {
				view.Lines = append(view.Lines, line)
			}
		}

		for _, header := range headers {
			if strings.HasPrefix(view.Header, "<"+header) {
				view.Header = strings.Replace(view.Header, "<", "", 1)
				views = append(views, view)
				break
			}
		}
	}

	return views
}

func ParseProperties(views []View2) []View2 {

	for index, view := range views {
		for _, line := range view.Lines {

			var property Property

			before, after, found := strings.Cut(line, "=")
			if found {
				property.Field = before
				property.Value = extractFirstString(after)
				view.Properties = append(view.Properties, property)
			}
		}
		views[index] = view
	}
	for _, view := range views {
		fmt.Println(view.Header, view.Properties)
	}
	return views
}

func EndParseViews(views []View2) {

	fmt.Print("\n\n")
	fmt.Println("=================================")
	fmt.Println("EndParseViews")

	var jsonString string

	for _, view := range views {

		jsonString = "{"
		for _, property := range view.Properties {

			if strings.Contains(property.Field, "Array") {
				fmt.Println(property.Value)
				strSlice, err := ParsArray(property.Value)
				if err == nil {
					// Convert slice of strings to JSON
					jsonData, err := json.Marshal(strSlice)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					// Convert JSON data to a string
					jsonStr := string(jsonData)
					jsonString += "\"" + property.Field + "\":" + jsonStr + ","
				}
			} else if strings.Contains(property.Field, "Color") || strings.Contains(property.Field, "color") {
				var color = ParseColor(property.Value)
				jsonString += "\"" + property.Field + "\":" + strconv.Itoa(color) + ","
			} else if strings.Contains(property.Value, "\"") {
				jsonString += "\"" + property.Field + "\":" + property.Value + ","
			} else if strings.Contains(property.Value, "'") {
				property.Value = strings.ReplaceAll(property.Value, "'", "")
				jsonString += "\"" + property.Field + "\":" + "\"" + property.Value + "\","
			} else {
				jsonString += "\"" + property.Field + "\":" + property.Value + ","
			}
		}

		jsonString = jsonString[:len(jsonString)-1]
		jsonString += "}"
		model.AddView(view.Header, jsonString)
	}

	fmt.Println("=================================")
	fmt.Print("\n\n")
}

//type ViewParse struct {
//	ViewName string
//	Fields   []string
//	Values   []string
//}

//func parseViewBase(viewParse ViewParse) model.ViewBase {
//	var dto model.ViewBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//
//		var value = viewParse.Values[index]
//
//		switch field {
//		case "id":
//			dto.Id = ParseString(value)
//			break
//		case "dx":
//			dto.Dx = ParseFloat(value)
//			break
//		case "dy":
//			dto.Dy = ParseFloat(value)
//			break
//		case "width":
//			dto.Width = ParseFloat(value)
//			break
//		case "height":
//			dto.Height = ParseFloat(value)
//			break
//		case "round":
//			dto.Round = ParseFloat(value)
//			break
//		case "padding":
//			dto.Padding = ParseFloat(value)
//			break
//		case "margin":
//			dto.Margin = ParseFloat(value)
//			break
//		case "icon":
//			dto.Icon = ParseString(value)
//		}
//		index++
//	}
//	return dto
//}
//
//func parseTextBase(viewParse ViewParse) model.TextBase {
//	var dto model.TextBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "text":
//			dto.Text = ParseString(value)
//			break
//		case "textColor":
//			dto.TextColor = ParseColor(value)
//			break
//		case "textSize":
//			dto.TextSize = ParseFloat(value)
//			break
//		case "align":
//			dto.Align = ParseString(value)
//			break
//		}
//		index++
//	}
//
//	return dto
//}
//
//func ParserThem(viewParse ViewParse) model.ThemBase {
//	var dto model.ThemBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "line":
//			dto.Line = ParseBoolean(value)
//			break
//		case "backgroundColor":
//			dto.BackgroundColor = ParseColor(value)
//			break
//		case "activeColor":
//			dto.ActiveColor = ParseColor(value)
//			break
//		case "circleColor":
//			dto.CircleColor = ParseColor(value)
//			break
//			break
//		}
//		index++
//	}
//	return dto
//}
//
//func ParserEventBase(viewParse ViewParse) model.EventBase {
//	var dto model.EventBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "onPress":
//			dto.OnPress = value
//			break
//		case "onClick":
//			dto.OnClick = value
//			break
//		}
//		index++
//	}
//	return dto
//}
//
////-----------------------------------------
//
//func TextParser(viewParse ViewParse) {
//	var dto model.TextView
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.TextBase = parseTextBase(viewParse)
//	textViews = append(textViews, dto)
//}
//
//func TextBoxParser(viewParse ViewParse) {
//	var dto model.TextBox
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.TextBase = parseTextBase(viewParse)
//	var index = 0
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "boxColor":
//			dto.BoxColor = ParseColor(value)
//			break
//		}
//		index++
//	}
//	textBoxes = append(textBoxes, dto)
//}
//
//func ImageParser(viewParse ViewParse) {
//	var dto model.Image
//	dto.ViewBase = parseViewBase(viewParse)
//	var index = 0
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "source":
//			dto.Source = ParseString(value)
//			break
//		}
//		index++
//	}
//
//	fmt.Println(dto)
//	images = append(images, dto)
//}
//
//func SwitchButtonParser(viewParse ViewParse) {
//	var dto model.SwitchButton
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "checked":
//			dto.Checked = ParseBoolean(value)
//			break
//		case "duration":
//			dto.Duration = ParseInteger(value)
//			break
//		}
//		index++
//	}
//	switchButtons = append(switchButtons, dto)
//}
//
//func CircleButtonParser(viewParse ViewParse) {
//	var dto model.CircleButton
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//	circleButtons = append(circleButtons, dto)
//}
//
//func SliderViewParser(viewParse ViewParse) {
//	var dto model.SliderView
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "circleSize":
//			dto.CircleSize = ParseFloat(value)
//			break
//		}
//		index++
//	}
//	sliderViews = append(sliderViews, dto)
//}
//
//func ChartViewParser(viewParse ViewParse) {
//	var dto model.ChartView
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//	var index = 0
//
//	StartArrayParse()
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//
//		switch field {
//		case "headerHeight":
//			dto.HeaderHeight = ParseFloat(value)
//			break
//		case "chartHeight":
//			dto.ChartHeight = ParseFloat(value)
//			break
//		case "footerHeight":
//			dto.FooterHeight = ParseFloat(value)
//			break
//		case "avatar":
//			dto.Avatar = ParseString(value)
//			break
//		case "title":
//			dto.Title = ParseString(value)
//			break
//		case "caption":
//			dto.Caption = ParseString(value)
//			break
//			//case "rowArray":
//			//	dto.RowArray, _ = ParsArray(value)
//			//	break
//			//case "columnArray":
//			//	dto.ColumnArray, _ = ParsArray(value)
//			//	break
//		}
//		index++
//	}
//
//	chartViews = append(chartViews, dto)
//}

//
//func parseNew(viewParse ViewParse) {
//
//	switch viewParse.Header {
//	case "SliderView":
//		aa(reflect.TypeOf(model.SliderView{}), viewParse)
//		break
//	case "ChartView":
//		aa(reflect.TypeOf(model.ChartView{}), viewParse)
//		break
//	}
//}
//
//func aa(t reflect.Type, viewParse ViewParse) {
//
//	view := reflect.New(t).Elem()
//
//	// Iterate over the fields of the struct
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		fmt.Printf("Field Name: %s, ", field.Name)
//
//		// Switch statement to identify the type of the field
//		switch field.Type.Kind() {
//		case reflect.String:
//			for index, viewParseField := range viewParse.Fields {
//				if strings.Compare(field.Name, viewParseField) == 0 {
//					var value = viewParse.Values[index]
//					view.FieldByName(field.Name).SetString(ParseString(value))
//				}
//				index++
//			}
//			break
//		case reflect.Int:
//			for index, viewParseField := range viewParse.Fields {
//				if strings.Compare(field.Name, viewParseField) == 0 {
//					var value = viewParse.Values[index]
//					view.FieldByName(field.Name).SetInt(int64(ParseInteger(value)))
//				}
//				index++
//			}
//			break
//		case reflect.Float32:
//			for index, viewParseField := range viewParse.Fields {
//				if strings.Compare(field.Name, viewParseField) == 0 {
//					var value = viewParse.Values[index]
//					view.FieldByName(field.Name).SetFloat(float64(ParseFloat(value)))
//				}
//				index++
//			}
//			break
//		case reflect.Slice:
//			if field.Type.Elem().Kind() == reflect.String {
//				//fmt.Println("Type--> []string")
//				for index, viewParseField := range viewParse.Fields {
//					if strings.Compare(field.Name, viewParseField) == 0 {
//						var value = viewParse.Values[index]
//						// Set the field of type string array dynamically
//						aliases, _ := ParsArray(value)
//						aliasesValue := reflect.ValueOf(aliases)
//						view.FieldByName(field.Name).Set(aliasesValue)
//					}
//					index++
//				}
//			} else {
//				fmt.Println("Type: slice of unknown type")
//			}
//			break
//		default:
//			fmt.Println("Field Type: unknown")
//		}
//	}
//
//	//chartViews = append(chartViews, view)
//}
