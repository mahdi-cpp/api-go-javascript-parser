package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"regexp"
	"strconv"
	"strings"
)

type Property struct {
	Field string
	Value string
}

type View struct {
	Id         string
	ViewName   string
	Variable   string
	Properties []Property
}

type Function struct {
	FuncName string
	Lines    []string
	Views    []View
}

type ViewDTO struct {
	Id       string `json:"id"`       // id for access to object
	ViewName string `json:"viewName"` // For example: 'Image' , 'SwitchButton' and ...
	Props    string `json:"props"`    // Properties of object
}

type FunctionDTO struct {
	FuncName string    `json:"funcName"`
	Views    []ViewDTO `json:"views"`
}

var functions []Function

func TestFunction() {
	functions = []Function{}
	_, err := ParseFunctions("web/index.js")
	if err != nil {
		return
	}
}

func ParseFunctions(fileName string) ([]FunctionDTO, error) {
	jsonString, err := readJsonFile(fileName)
	if err != nil {
		return nil, nil
	}

	jsonString = RemoveMultiLineComments(jsonString)
	jsonString = RemoveLineComments(jsonString)

	FindFunctions(jsonString)
	ParseViewsAndVariables()
	ParseViewProperties()
	functions := EndProcess()
	return functions, nil
}

func FindFunctions(jsonString string) []Function {
	fmt.Println("=======================================================")
	lines := strings.Split(jsonString, "\n")
	var index = 0

	for i := 0; i < len(lines); i++ {
		if len(lines) > index {
			if strings.HasPrefix(lines[index], "function") {
				var function Function
				funcName, err := getStringBetween(lines[index], "function", "(")

				if err != nil {
					fmt.Println("FindFunctions error")
				} else {
					funcName = strings.ReplaceAll(funcName, " ", "")
					function.FuncName = funcName
					for j := index + 1; j < len(lines); j++ {
						if strings.Contains(lines[j], "}") {
							functions = append(functions, function)
							break
						} else {
							fmt.Println(lines[j])
							function.Lines = append(function.Lines, lines[j])
						}
						index++
					}
				}
			}
			index++
		}
	}
	return functions
}

func ParseViewsAndVariables() {
	fmt.Println("-------------------------------")
	fmt.Println("ParseViewsAndVariables")

	// Create a regular expression pattern to match "let" as a whole word
	pattern := `\blet\b`

	//var newFunctions []Function

	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)
	var index = 0
	for _, function := range functions {
		for _, line := range function.Lines {
			if strings.HasPrefix(line, "let") {

				if re.MatchString(line) {
					var view View
					before, after, found := strings.Cut(line, "=")
					if found {
						before = strings.Replace(before, "let", "", 1)
						before = strings.ReplaceAll(before, " ", "")
						view.Variable = before

						view1, _, _ := strings.Cut(after, ";")
						view1, _, _ = strings.Cut(view1, "//")
						view1 = strings.Replace(view1, "Values", "", 1)
						view1 = strings.ReplaceAll(view1, " ", "")

						view.ViewName = view1
						//fmt.Println("---->", view.ViewName, view.Variable)
					}
					function.Views = append(function.Views, view)
				}
			}
		}
		functions[index].Views = append(functions[index].Views, function.Views...)
		index++
		//functions = append(functions, function)
	}
	fmt.Println("==============================")
	for _, function := range functions {
		fmt.Println("----------------------------------")
		fmt.Println(function.FuncName)
		for _, view := range function.Views {
			fmt.Println("         -->", view.ViewName, "      ", view.Variable)
		}
	}
	fmt.Println("==============================")
}

func ParseViewProperties() {

	fmt.Println("==============================")
	fmt.Println("ParseViewProperties")

	var index = 0
	for _, function := range functions {
		fmt.Println("-------------------")
		var views []View
		for _, view := range function.Views {
			for _, line := range function.Lines {
				if strings.HasPrefix(line, view.Variable+".") {
					//fmt.Println(line)

					var property Property

					_, after, found := strings.Cut(line, view.Variable+".")
					if found {
						b, a, _ := strings.Cut(after, "=")

						a, _, _ = strings.Cut(a, ";")
						a, _, _ = strings.Cut(a, "//")
						a = strings.ReplaceAll(a, " ", "")
						//a = strings.ReplaceAll(a, "'", "")

						b = strings.ReplaceAll(b, " ", "")

						property.Field = b
						property.Value = a
						view.Properties = append(view.Properties, property)

					} else {
						fmt.Println("not found", view.Variable+".")
					}
				}
			}
			views = append(views, view)
		}
		function.Views = views
		fmt.Println(function.Views)
		functions[index].Views = function.Views
		index++
	}

	fmt.Println("==============================")
}

func EndProcess() []FunctionDTO {

	var functionsDTO []FunctionDTO

	for _, function := range functions {
		var functionDTO FunctionDTO
		functionDTO.FuncName = function.FuncName

		for _, view := range function.Views {
			var viewDTO ViewDTO
			viewDTO.ViewName = view.ViewName
			viewDTO.Props = "{"
			for _, property := range view.Properties {
				if strings.Compare(property.Field, "id") == 0 {
					property.Value = strings.ReplaceAll(property.Value, "'", "")
					viewDTO.Id = property.Value
				} else if strings.Contains(property.Field, "Color") {
					var color = utils.GetColor(property.Value)
					viewDTO.Props += "\"" + property.Field + "\":" + strconv.Itoa(color) + ","
				} else if strings.Contains(property.Value, "'") { // if is values string need double quotation
					property.Value = strings.ReplaceAll(property.Value, "'", "")
					viewDTO.Props += "\"" + property.Field + "\":" + "\"" + property.Value + "\","
				} else {
					viewDTO.Props += "\"" + property.Field + "\":" + property.Value + ","
				}
			}
			viewDTO.Props = viewDTO.Props[:len(viewDTO.Props)-1]
			viewDTO.Props += "}"
			functionDTO.Views = append(functionDTO.Views, viewDTO)
		}

		functionsDTO = append(functionsDTO, functionDTO)
	}

	return functionsDTO
}

// RestFunctions -------------------------------------------
func RestFunctions() []FunctionDTO {
	functions = []Function{}
	f, _ := ParseFunctions("web/index.js")
	return f
}
