package repository

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func ParseBoolean(str string) bool {
	if strings.Contains(str, "true") {
		return true
	}
	return false
}

//func ParseInteger(str string) int {
//	i, err := strconv.Atoi(str)
//	if err != nil {
//		return 0
//	}
//	return i
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
