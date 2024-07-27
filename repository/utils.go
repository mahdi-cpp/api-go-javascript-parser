package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"strconv"
	"strings"
)

func RemoveLineComments(str string) string {

	lines := strings.Split(str, "\n")
	//fmt.Println("=======================================================")
	var newString = ""
	for _, line := range lines {

		line = removeLeadingSpaces(line)

		if isOnlySpaces(line) {
			//fmt.Println("******isOnlySpaces*********")
		} else if startLineComment(line) {
			//fmt.Println("find_line_comment *********")
		} else {
			//line = removeLeadingSpaces(line)
			newString += line + "\n"
			//fmt.Println(line)
			//fmt.Println("--------------------------------------------------")
		}

	}
	//fmt.Println("=======================================================")
	return newString
}

func RemoveMultiLineComments(str string) string {

	//var comments []comment

	for {
		startIndex := 0
		endIndex := 0

		startIndex = strings.Index(str, "/*")
		if startIndex == -1 {
			break
		} else {
			endIndex = strings.Index(str, "*/")
			if endIndex == -1 {
				break
			} else {
				str = removeStartEndByIndex(str, startIndex, endIndex+1)
			}
		}
	}

	return str

	//for _, comment := range comments {
	//	removeStartEndByIndex(str, int(comment.start), int(comment.end))
	//}

}

func startsWithSpace(s string) bool {
	return strings.HasPrefix(s, " ")
}

func startsWithTab(s string) bool {
	return strings.HasPrefix(s, "\t")
}

func removeStartEndByIndex(s string, startIndex, endIndex int) string {
	if startIndex < 0 || endIndex < 0 || startIndex >= len(s) || endIndex >= len(s) || startIndex > endIndex {
		return "Invalid index positions"
	}

	return s[:startIndex] + s[endIndex+1:]
}

func isOnlySpaces(s string) bool {
	return strings.TrimSpace(s) == ""
}

func removeLeadingSpaces(s string) string {
	return strings.TrimLeft(s, " ")
}

func startLineComment(s string) bool {
	return strings.HasPrefix(s, "//")
}

func getStringBetween(s, start, end string) (string, error) {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return "", fmt.Errorf("start string '%s' not found", start)
	}

	endIndex := strings.Index(s[startIndex+len(start):], end)
	if endIndex == -1 {
		return "", fmt.Errorf("end string '%s' not found after start string", end)
	}

	return s[startIndex+len(start) : startIndex+len(start)+endIndex], nil
}

func readJsonFile(file string) (string, error) {
	jsFile, err := utils.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	} else {
		fmt.Println("Ok reading file: '" + file + "'")
	}

	return jsFile, nil
}

// Method to remove an element from the Properties slice by index
func (v *View) removePropertyByIndex(index int) {
	if index < 0 || index >= len(v.Properties) {
		fmt.Println("Index out of range")
		return
	}

	v.Properties = append(v.Properties[:index], v.Properties[index+1:]...)
}

func ParseString(str string) string {
	//str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")
	return str
}

func ParseInteger(str string) int {
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		// do something sensible
	}
	return int(value)
}

func ParseFloat(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		// do something sensible
	}
	return float32(value)
}

func ParseColor(strColor string) int {
	strColor = strings.Replace(strColor, "'", "", 2)
	color, err := utils.Parse(strColor)
	if err == nil {
		var rgb = int(color.ToRGB().R)
		rgb = (rgb << 8) + int(color.ToRGB().G)
		rgb = (rgb << 8) + int(color.ToRGB().B)
		return rgb
	}

	fmt.Println(err)
	return 0
}
