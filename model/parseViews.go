package model

import (
	"encoding/json"
	"fmt"
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
	Clear()
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

			if strings.Contains(property.Field, "onClick") {
				jsonString += "\"" + property.Field + "\":" + "\"" + property.Value + "\","
			} else if strings.Contains(property.Field, "dataArray") {
				fmt.Println(property.Value)
				intArray, err := ParseFloatArray(property.Value)
				if err == nil {
					// Convert slice of strings to JSON
					jsonData, err := json.Marshal(intArray)
					if err != nil {
						fmt.Println("Error 5:", err)
						return
					}
					// Convert JSON data to a string
					jsonStr := string(jsonData)
					fmt.Println("==================================")
					fmt.Println(intArray)
					fmt.Println("==================================")
					jsonString += "\"" + property.Field + "\":" + jsonStr + ","
					fmt.Println(jsonString)
				}
			} else if strings.Contains(property.Field, "Array") {
				fmt.Println(property.Value)
				strSlice, err := ParsArray(property.Value)
				if err == nil {
					// Convert slice of strings to JSON
					jsonData, err := json.Marshal(strSlice)
					if err != nil {
						fmt.Println("Error 7:", err)
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
		AddView(view.Header, jsonString)
	}

	fmt.Println("=================================")
	fmt.Print("\n\n")
}
