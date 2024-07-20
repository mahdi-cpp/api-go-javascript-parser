package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"strings"
)

var images []model.Image
var textViews []model.TextView
var circleButtons []model.CircleButton
var switchButtons []model.SwitchButton

func GetTextViews() []model.TextView {
	return textViews
}
func GetImages() []model.Image {
	return images
}
func GetCircleButtons() []model.CircleButton {
	return circleButtons
}
func GetSwitchButtons() []model.SwitchButton {
	return switchButtons
}

func StartScriptParse() {
	images = []model.Image{} //set to Empty Array
	textViews = []model.TextView{}
	circleButtons = []model.CircleButton{}
	switchButtons = []model.SwitchButton{}

	scriptParse("web/ViewSlider.js")
}

func scriptParse(file string) {

	jsFile, err := utils.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	} else {
		fmt.Println("Ok reading file: '" + file + "'")
	}

	var search = "ViewSlider"
	_, after, found := strings.Cut(jsFile, search)

	if found {
		var insideSearch = "</View>"
		before, _, hasFloatView := strings.Cut(after, insideSearch)
		if hasFloatView {
			FindView(before)
		} else {
			fmt.Println("Not Found: " + "'" + search + "'")
		}
	} else {
		fmt.Println("Not Found: " + "'" + search + "'")
	}
}

func FindView(view string) {

	fmt.Println("-------------------------------")
	fmt.Println(view)

	var search = "<View>"
	_, after, found := strings.Cut(view, search)

	if found {
		//var search2 = "</View>"
		//before, _, hasView := strings.Cut(after, search2)
		//if hasView {
		elements := strings.Split(after, "<")
		for _, element := range elements {
			if hasElement(element) {
				ViewElementParse(element)
			}
		}
		//} else {
		//	fmt.Println("Not Found: " + "'" + search2 + "'")
		//}
	} else {
		fmt.Println("Not Found: " + "'" + search + "'")
	}
}

type ViewParse struct {
	ViewName   string
	Attributes []string
	Values     []string
}

func ViewElementParse(view string) {

	attributes := splitWithStartAlphabetWordExceptCurlyBraces(view)
	//fmt.Println(view)
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
				viewParse.Attributes = append(viewParse.Attributes, attribute)
			}
		}
		index += 1
	}

	fmt.Println(viewParse.ViewName)
	//fmt.Println(viewParse.Attributes)
	//fmt.Println(viewParse.Values)

	switch viewParse.ViewName {
	case "TextView":
		TextParser(viewParse)
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
	}

	fmt.Println("--------------------------------------------------")

}
