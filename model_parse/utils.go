package model_parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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

func prettyJsonConsole(jsonStr string) string {

	jsonStr2 := `{
		"chartViews": "{\n  \"avatar\": \"call2/ali5.jpg\",\n  \"backgroundColor\": 16750592,\n  \"chartHeight\": 250,\n  \"columnArray\": [\n    \"100\",\n    \"90\",\n    \"80\",\n    \"70\",\n    \"60\",\n    \"50\",\n    \"40\",\n    \"30\",\n    \"20\",\n    \"10\"\n  ],\n  \"dx\": 10,\n  \"dy\": 900,\n  \"footerHeight\": 86,\n  \"headerHeight\": 60,\n  \"icon\": \"icons/health_blood_pressure.png\",\n  \"id\": \"Mahdi 01236\",\n  \"margin\": 0,\n  \"padding\": 25,\n  \"round\": 0,\n  \"rowArray\": [\n    \"Mahdi\",\n    \"Ali\",\n    \"Reza\",\n    \"Sara\",\n    \"Maryam\",\n    \"Mahyar\",\n    \"Parsa\"\n  ],\n  \"title\": \"Blood pressure\",\n  \"width\": 450\n}"
	}`

	// Unmarshal the JSON string into an interface{}
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonStr2), &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Marshal the interface{} back to JSON with indentation for pretty printing
	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Convert the pretty-printed JSON to a string
	prettyStr := string(prettyJSON)

	// Print the pretty-printed JSON string
	fmt.Println("Formatted JSON:")
	fmt.Println(prettyStr)

	return prettyStr
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

func extractFirstString(input string) string {
	// Define a regular expression to match the first string inside curly braces '{}'
	re := regexp.MustCompile(`\{([^}]+)\}`)

	// Find the first submatch in the input string
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		return match[1]
	}

	return ""
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

func hasValidModel(view string) (bool, error) {

	// Open the file for reading
	file, err := os.Open("web/models.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the lines
	var lines []string

	// Read lines from the file and store them in the slice
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	// Print the lines read from the file
	for _, line := range lines {
		if strings.Compare(line, view) == 0 {
			return true, nil
		}
	}

	return false, nil
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
