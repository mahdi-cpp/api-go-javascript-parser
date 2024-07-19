package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-movie/javascript/models"
	"github.com/mahdi-cpp/api-go-movie/utils"
	"regexp"
	"strings"
	"unicode"
)

var images []models.Image

func GetImages() []models.Image {
	images := ScriptParse()
	return images
}

func ScriptParse() []models.Image {

	jsFile, err := utils.ReadFile("web/FloatView.js")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	} else {
		fmt.Println("ok reading file")
	}

	_, after, found := strings.Cut(jsFile, "<FloatView")

	if found {
		FloatView, _, hasFloatView := strings.Cut(after, " </FloatView>")
		if hasFloatView {
			fmt.Println(FloatView)
			FindView(FloatView)
		} else {
			fmt.Println("FloatView Not Exist")
		}
	}

	return images
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
					//fmt.Println(element)
					ViewElementParse(element)
					//fmt.Println("-------------------------------")
				}
			}
		} else {
			fmt.Println("View Not Exist")
		}
	}
}

func ViewElementParse(element string) {

	//attributes := strings.Split(element, " ")

	//str := "title={'Thetitle and new helium handler are required. It is recommended to'}"
	attributes := splitWithStartAlphabetWordExceptCurlyBraces(element)

	var attrs []string
	var values []string

	for _, attribute := range attributes {

		if strings.Contains(attribute, "{") {
			//attribute = strings.Replace(attribute, "{", "", -1)
			//attribute = strings.Replace(attribute, "}", "", -1)
			values = append(values, attribute)
		} else {
			attrs = append(attrs, attribute)
		}
	}

	fmt.Println("Element:", attrs[0])

	if strings.Contains(attrs[0], "Image") {
		//var img javascript.Image
		//img.Source = ""
		//src, exists :=  AttrFloat("width")
		//if exists {
		//	//img.Width = src
		//}
		//images = append(images, img)
	}

	attrs = removeFirstElement(attrs)
	fmt.Println("Attribute:", attrs)
	fmt.Println("Values:", values)

	fmt.Println("--------------------------------------------------")

}

//type ScriptParser struct {
//	images          []model.Image
//	children        map[string]model.Image
//	register        chan *Client
//	unregister      chan *Client
//	rooms           map[*Room]bool
//	imageRepository model.ImageRepository
//}

//func AttrFloat(attrName string) (val string, exists bool) {
//
//	return nil, nil
//}

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
