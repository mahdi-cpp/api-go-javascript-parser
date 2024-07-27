package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"regexp"
	"strings"
)

var textViews []model.TextView
var textBoxes []model.TextBox
var images []model.Image
var circleButtons []model.CircleButton
var switchButtons []model.SwitchButton
var sliderViews []model.SliderView
var chartViews []model.ChartView

func RestTextBoxes() []model.TextBox {
	return textBoxes
}
func RestTextViews() []model.TextView {
	return textViews
}
func RestImages() []model.Image {
	return images
}
func RestCircleButtons() []model.CircleButton {
	return circleButtons
}
func RestSwitchButtons() []model.SwitchButton {
	return switchButtons
}
func RestSliderView() []model.SliderView {
	return sliderViews
}
func RestChartViews() []model.ChartView {
	return chartViews
}

func StartScriptParse() {
	textBoxes = []model.TextBox{}
	textViews = []model.TextView{}
	images = []model.Image{} //set to Empty Array
	circleButtons = []model.CircleButton{}
	switchButtons = []model.SwitchButton{}
	sliderViews = []model.SliderView{}
	chartViews = []model.ChartView{}

	TestParseViews()
}

type ViewParse struct {
	ViewName string
	Fields   []string
	Values   []string
}

type View2 struct {
	Header     string
	Properties []Property
	Lines      []string
}

func TestParseViews() {
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
		before, _, hasFloatView := strings.Cut(after, insideSearch)
		if hasFloatView {
			ParseViews(before)
		} else {
			fmt.Println("Not Found: " + "'" + search + "'")
		}
	} else {
		fmt.Println("Not Found: " + "'" + search + "'")
	}
}

func ParseViews(jsonViews string) {

	fmt.Println("============================")
	fmt.Println("ParseViews")

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
				views = append(views, view)
				break
			}
		}
	}

	fmt.Println("============================")
	for _, view := range views {
		fmt.Println(view.Header)
		for _, line := range view.Lines {
			fmt.Println("         -->", line)
		}
		fmt.Println("------------------------")
	}

	fmt.Println("============================\n")
}

func ViewElementParse(view string) {

	attributes := splitWithStartAlphabetWordExceptCurlyBraces(view)
	//fmt.Println(View)
	//fmt.Println(attributes)

	var viewParse ViewParse
	var index = 0
	for _, attribute := range attributes {

		if strings.Contains(attribute, "{") {
			attribute = strings.Replace(attribute, "{", "", -1)
			attribute = strings.Replace(attribute, "}", "", -1)
			viewParse.Values = append(viewParse.Values, attribute)
		} else {
			if index == 0 {
				viewParse.ViewName = attribute
			} else {
				attribute = strings.Replace(attribute, " ", "", 2)
				viewParse.Fields = append(viewParse.Fields, attribute)
			}
		}
		index += 1
	}

	fmt.Println("============================================")
	fmt.Println(viewParse.ViewName)
	fmt.Println(viewParse.Fields)
	fmt.Println(viewParse.Values)
	fmt.Println("============================================")
	//fmt.Println(viewParse.Fields)
	//fmt.Println(viewParse.Values)

	switch viewParse.ViewName {
	case "TextView":
		TextParser(viewParse)
		break
	case "TextBox":
		TextBoxParser(viewParse)
		break
	case "Image":
		ImageParser(viewParse)
		break
	case "CircleButton":
		CircleButtonParser(viewParse)
		break
	case "SwitchButton":
		SwitchButtonParser(viewParse)
		break
	case "SliderView":
		SliderViewParser(viewParse)
		break
	case "ChartView":
		ChartViewParser(viewParse)
		break
	}

	fmt.Println("--------------------------------------------------")

}

func splitStringWithHeaders(input string, headers []string) []string {
	var result []string

	lines := strings.Split(input, "\n")

	var currentSection strings.Builder
	var inSection bool

	for _, line := range lines {

		for _, header := range headers {
			if strings.HasPrefix(line, "<"+header) {
				if inSection {
					result = append(result, currentSection.String())
					currentSection.Reset()
				}
				inSection = true
			}
		}

		if inSection {
			currentSection.WriteString(line + "\n")
		}
	}

	if inSection {
		result = append(result, currentSection.String())
	}

	return result
}
