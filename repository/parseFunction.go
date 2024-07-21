package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"strconv"
	"strings"
)

type View struct {
	Id    string `json:"id"`    // id for access to object
	View  string `json:"view"`  // For example: 'Image' , 'SwitchButton' and ...
	Param string `json:"param"` // properties of object
}

type Function struct {
	FuncName string `json:"funcName"`
	Views    []View `json:"objects"`
}

var functions []Function

func StartFunction() {
	functions = []Function{}
	functionParse("web/index.js")
}
func GetFunctions() []Function {
	return functions
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

type funcViewParse struct {
	Id       string
	View     string
	Param    string
	Variable string
}

func findFuncLets(lines []string) {

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("findFuncLets")

	var views []funcViewParse
	var funcView View
	var funcViews []View
	var function Function

	//for _, line := range lines {
	//	if strings.HasPrefix(line, "let") {
	//		println("'" + line + "'")
	//	}
	//}
	fmt.Println("-------------------------------")

	var lets []string
	//var variables []string
	//var objs []string

	for _, line := range lines {
		if strings.HasPrefix(line, "let") {
			lets = append(lets, line)
		}
	}

	for _, let := range lets {
		var view funcViewParse
		before, after, found := strings.Cut(let, "=")

		if found {
			before = strings.Replace(before, "let", "", 1)
			before = strings.ReplaceAll(before, " ", "")
			view.Variable = before

			b, _, _ := strings.Cut(after, ";")
			b = strings.Replace(b, "Values", "", 1)
			b = strings.ReplaceAll(b, " ", "")
			view.View = b
		}
		views = append(views, view)
	}

	//for _, view := range views {
	//	fmt.Println(view.View)
	//}
	//
	//for _, line := range lines {
	//	fmt.Println(line)
	//}

	for _, view := range views { //values
		view.Param = "{"

		for _, line := range lines {
			if !strings.HasPrefix(line, "let") {

				if strings.Contains(line, view.Variable) {
					before, after, found := strings.Cut(line, "=")
					if found {
						before = strings.Replace(before, view.Variable, "", 1)
						before = strings.Replace(before, ".", "", 1)
						before = strings.ReplaceAll(before, " ", "")

						value, _, _ := strings.Cut(after, ";")
						value = strings.ReplaceAll(value, " ", "")

						if strings.Compare(before, "id") == 0 {
							value = strings.ReplaceAll(value, "'", "")
							view.Id = value
						} else if strings.Contains(before, "Color") {
							value = strings.ReplaceAll(value, "\"", "")
							var color = utils.GetColor(value)
							view.Param += "\"" + before + "\":" + strconv.Itoa(color) + ","
						} else if strings.Contains(value, "'") { // if is values string need double quotation
							value = strings.ReplaceAll(value, "'", "")
							view.Param += "\"" + before + "\":" + "\"" + value + "\","
						} else {
							view.Param += "\"" + before + "\":" + value + ","
						}
					}
				}
			}
		}
		view.Param = view.Param[:len(view.Param)-1]
		view.Param += "}"

		fmt.Println(view.Id)
		fmt.Println(view.View)
		fmt.Println(view.Param)

		funcView.Id = view.Id
		funcView.View = view.View
		funcView.Param = view.Param

		funcViews = append(funcViews, funcView)

		fmt.Println("------------------------------")
	}

	function.FuncName = "changePhoto"
	function.Views = funcViews
	functions = append(functions, function)

	fmt.Println(function)
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
