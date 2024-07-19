package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var images []model.Image
var textViews []model.TextView
var circleButtons []model.CircleButton
var switchButtons []model.SwitchButton

//-------------------------------------------------------------------

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

//-------------------------------------------------------------------

func StartScriptParse() {
	images = []model.Image{} //set to Empty Array
	textViews = []model.TextView{}
	circleButtons = []model.CircleButton{}
	switchButtons = []model.SwitchButton{}
	scriptParse()
}

func scriptParse() {

	jsFile, err := utils.ReadFile("web/FloatView.js")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	} else {
		fmt.Println("ok reading file")
	}

	_, after, found := strings.Cut(jsFile, "<FloatView")

	if found {
		FloatView, _, hasFloatView := strings.Cut(after, " </FloatView>")
		if hasFloatView {
			//fmt.Println(FloatView)
			FindView(FloatView)
		} else {
			fmt.Println("FloatView Not Exist")
		}
	}
}

func FindView(FloatView string) {

	fmt.Println("-------------------------------")

	_, after, found := strings.Cut(FloatView, "<View>")

	if found {
		View, _, hasView := strings.Cut(after, "</View>")
		if hasView {
			elements := strings.Split(View, "<")
			for _, element := range elements {
				if hasElement(element) {
					fmt.Println(element)
					ViewElementParse(element)
					//fmt.Println("-------------------------------")
				}
			}
		} else {
			fmt.Println("View Not Exist")
		}
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

func TextParser(view ViewParse) {

	var textView model.TextView
	var index = 0
	for _, attr := range view.Attributes {
		//fmt.Println("'" + attr + "'")
		switch attr {
		case "dx":
			textView.Dx = utils.GetFloat(view.Values[index])
			break
		case "dy":
			textView.Dy = utils.GetFloat(view.Values[index])
			break
		case "width":
			textView.Width = utils.GetFloat(view.Values[index])
			break
		case "height":
			textView.Height = utils.GetFloat(view.Values[index])
			break
		case "text":
			var t = view.Values[index]
			t = strings.Replace(t, "'", "", 2)
			textView.Text = t
			break
		case "textColor":
			textView.TextColor = utils.GetColor(view.Values[index])
			break
		case "textSize":
			textView.TextSize = utils.GetFloat(view.Values[index])
			break
		case "align":
			var t = view.Values[index]
			t = strings.Replace(t, "'", "", 2)
			textView.Align = t
			break
		}
		index++
	}

	textViews = append(textViews, textView)
}

func ImageParser(view ViewParse) {
	var image model.Image
	var index = 0
	for _, attr := range view.Attributes {
		switch attr {
		case "dx":
			image.Dx = utils.GetFloat(view.Values[index])
			break
		case "dy":
			image.Dy = utils.GetFloat(view.Values[index])
			break
		case "width":
			image.Width = utils.GetFloat(view.Values[index])
			break
		case "source":
			var t = view.Values[index]
			t = strings.Replace(t, "'", "", 2)
			image.Source = t
			break
		case "round":
			image.Round = utils.GetFloat(view.Values[index])
			break
		case "height":
			image.Height = utils.GetFloat(view.Values[index])
			break
		}
		index++
	}

	fmt.Println(image)
	images = append(images, image)
}

func SwitchButtonParser(view ViewParse) {
	var switchButton model.SwitchButton
	var index = 0
	for _, attr := range view.Attributes {
		//fmt.Println("'" + attr + "'")
		switch attr {
		case "dx":
			switchButton.Dx = utils.GetFloat(view.Values[index])
			break
		case "dy":
			switchButton.Dy = utils.GetFloat(view.Values[index])
			break
		case "width":
			switchButton.Width = utils.GetFloat(view.Values[index])
			break
		case "height":
			switchButton.Height = utils.GetFloat(view.Values[index])
			break
		case "value":
			// we can also convert negative numbers
			s1, err1 := strconv.Atoi(view.Values[index])
			if err1 != nil {
				fmt.Println("Can't convert this to an int!")
			} else {
				switchButton.Value = s1
			}
			break
		}
		index++
	}

	switchButtons = append(switchButtons, switchButton)
}

func CircleButtonParser(view ViewParse) {

	var circleButton model.CircleButton

	fmt.Println("ImageParser:", view)
	var index = 0
	for _, attr := range view.Attributes {
		//fmt.Println("'" + attr + "'")
		switch attr {
		case "dx":
			circleButton.Dx = utils.GetFloat(view.Values[index])
			break
		case "dy":
			circleButton.Dy = utils.GetFloat(view.Values[index])
			break
		case "width":
			circleButton.Width = utils.GetFloat(view.Values[index])
			break
		case "height":
			circleButton.Height = utils.GetFloat(view.Values[index])
			break
		case "icon":
			var t = view.Values[index]
			t = strings.Replace(t, "'", "", 2)
			circleButton.Icon = t
			break
		}
		index++
	}

	circleButtons = append(circleButtons, circleButton)
}

//------------------------------------------------------

func removeFirstElement(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	return slice[1:]
}

func splitWithStartAlphabetWordExceptCurlyBraces(s string) []string {
	var result []string

	re := regexp.MustCompile(`{[^{}]+}|[a-zA-Z]\w+`)
	words := re.FindAllString(s, -1)

	for _, word := range words {
		if !strings.Contains(word, "{") && !strings.Contains(word, "}") {
			if len(word) > 0 {
				firstChar := []rune(word)[0]
				if unicode.IsLetter(firstChar) {
					result = append(result, word)
				}
			}
		} else {
			result = append(result, word)
		}
	}

	return result
}

func hasElement(element string) bool {

	if strings.HasPrefix(element, "Text") {
		return true
	}
	if strings.HasPrefix(element, "Image") {
		return true
	}
	if strings.HasPrefix(element, "CircleButton") {
		return true
	}
	if strings.HasPrefix(element, "SwitchButton") {
		return true
	}
	return false

}

func splitWithStartAlphabetWordExceptQuoted(s string) []string {
	var result []string

	re := regexp.MustCompile(`"[^"]+"|\S+`)
	words := re.FindAllString(s, -1)

	for _, word := range words {
		if !strings.Contains(word, "\"") {
			if len(word) > 0 {
				firstChar := []rune(word)[0]
				if unicode.IsLetter(firstChar) {
					result = append(result, word)
				}
			}
		} else {
			result = append(result, word)
		}
	}

	return result
}

func splitWithStartAlphabetWord(s string) []string {
	var result []string
	words := strings.Fields(s)

	for _, word := range words {
		if len(word) > 0 {
			firstChar := []rune(word)[0]
			if unicode.IsLetter(firstChar) {
				result = append(result, word)
			}
		}
	}

	return result
}

func startsWithAlphabet(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstChar := []rune(s)[0]
	return unicode.IsLetter(firstChar) && !unicode.IsSpace(firstChar)
}

func substringWithSearch(str, search string) string {
	index := strings.Index(str, search)
	if index == -1 {
		return "" // Search pattern not found
	}
	return str[index+len(search):]
}
