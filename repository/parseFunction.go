package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"strings"
)

type View struct {
	Id     string `json:"id"`     // id for access to object
	Object string `json:"object"` // For example: 'Image' , 'SwitchButton' and ...
	Param  string `json:"param"`  // properties of object
}

type Function struct {
	FuncName string `json:"funcName"`
	Views    []View `json:"objects"`
}

func StartFunction() {
	functionParse("web/index.js")
}

func functionParse(fileName string) {

	jsFile, err := readJsonFile(fileName)
	if err != nil {
		return
	}

	var search = "function"
	_, after, found := strings.Cut(jsFile, search)

	var funcName = ""

	if found {
		var insideSearch = "{"
		before, after, hasSearch := strings.Cut(after, insideSearch)
		if hasSearch {
			funcName = strings.Replace(before, " ", "", 2)
			funcName = strings.Replace(funcName, "(", "", 1)
			funcName = strings.Replace(funcName, ")", "", 1)
			fmt.Println("'" + funcName + "'")
			findFuncBody(after)
		} else {
			fmt.Println("Not Found: " + "'" + search + "'")
		}
	} else {
		fmt.Println("Not Found: " + "'" + search + "'")
	}
}

func findFuncBody(body string) {
	fmt.Println("-------------------------------")
	fmt.Println("findFuncBody")

	var search = "}"
	before, _, found := strings.Cut(body, search)

	if found {
		println(before)
	} else {
		fmt.Println("Not Found: " + "'" + search + "'")
	}

	lines := strings.Split(before, "\n")
	var outputLines []string

	for _, line := range lines {
		// Check if the line contains only spaces
		if strings.TrimSpace(line) != "" {
			line := removeLeadingSpaces(line)
			outputLines = append(outputLines, line)
		}
	}

	findFuncLets(outputLines)

	//fmt.Println("-------------------------------")
	//for _, str2 := range outputLines {
	//	fmt.Println(str2)
	//}

	fmt.Println("-------------------------------")
}

//let switch1 = SwitchButtonValues
//let switch2 = SwitchButtonValues
//let imageTest = ImageValues

type FuncViewParse struct {
	view     string
	variable string
	Param    string
}

func findFuncLets(lines []string) {

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("findFuncLets")

	for _, line := range lines {
		if strings.HasPrefix(line, "let") {
			println("'" + line + "'")
		}
	}
	fmt.Println("-------------------------------")

	var lets []string
	//var variables []string
	//var objs []string
	var views []FuncViewParse

	for _, line := range lines {
		if strings.HasPrefix(line, "let") {
			lets = append(lets, line)
		}
	}
	fmt.Println("------------------------------------")
	for _, let := range lets {
		var view FuncViewParse
		before, after, found := strings.Cut(let, "=")

		if found {
			before = strings.Replace(before, "let", "", 4)
			before = strings.Replace(before, " ", "", 4)
			view.variable = before

			after = strings.Replace(after, "Values", "", 1)
			after = strings.Replace(after, " ", "", 1)
			view.view = after
		}
		views = append(views, view)
	}

	for _, view := range views {
		fmt.Println(view.view)
		fmt.Println(view.variable)

		for _, line := range lines {
			if !strings.HasPrefix(line, "let") {

			}

		}

	}
}

// --------------------------------------------------------------------------

func removeLeadingSpaces(input string) string {
	for i, char := range input {
		if char != ' ' {
			return input[i:]
		}
	}
	return ""
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
