package main

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
)

func main() {

	model.StartArrayParse()
	//repository.TestFunction()

	model.StartParseViews()

	Run()

	//vm := goja.New()
	//data, err := os.ReadFile("web/data.js")
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Execute JavaScript code that defines and returns a constant array
	//_, err = vm.RunString(string(data))
	//if err != nil {
	//	fmt.Println("Error executing JavaScript:", err)
	//	return
	//}
	//
	//// Call the JavaScript function to get the constant array
	//result, err := vm.RunString("users")
	//if err != nil {
	//	fmt.Println("Error calling JavaScript function:", err)
	//	return
	//}
	//
	//// Convert the JavaScript array to a Go slice of integers
	//jsArray := result.Export().([]interface{})
	//goArray := make([]string, len(jsArray))
	//for i, v := range jsArray {
	//	goArray[i] = v.(string)
	//}
	//
	//for _, a := range jsArray {
	//	fmt.Println(a)
	//}

}
