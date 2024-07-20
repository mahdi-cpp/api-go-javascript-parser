package repository

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ParseBoolean(str string) bool {
	if strings.Contains(str, "true") {
		return true
	}
	return false
}

func ParseInteger(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
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
