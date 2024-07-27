package repository

import (
	"fmt"
	"github.com/dop251/goja"
	"os"
)

var vm2 *goja.Runtime

func StartArrayParse() {
	vm := goja.New()
	data, err := os.ReadFile("web/data.js")
	if err != nil {
		panic(err)
	}

	// Execute JavaScript code that defines and returns a constant array
	_, err = vm.RunString(string(data))
	if err != nil {
		fmt.Println("Error executing JavaScript:", err)
		return
	}

	vm2 = vm
}

func ParsArray(arrayName string) ([]string, error) {

	result, err := vm2.RunString(arrayName)
	if err != nil {
		fmt.Println("Error calling JavaScript array:", err)
		return nil, err
	}

	// Convert the JavaScript array to a Go slice of integers
	jsArray := result.Export().([]interface{})
	goArray := make([]string, len(jsArray))
	for i, v := range jsArray {
		goArray[i] = v.(string)
	}

	for _, a := range jsArray {
		fmt.Println(a)
	}
	return goArray, nil
}

func ParsArray2(arrayName string) ([]int64, error) {

	result, err := vm2.RunString(arrayName)
	if err != nil {
		fmt.Println("Error calling JavaScript array:", err)
		return nil, err
	}

	// Convert the JavaScript array to a Go slice of integers
	jsArray := result.Export().([]interface{})
	goArray := make([]int64, len(jsArray))
	for i, v := range jsArray {
		goArray[i] = v.(int64)
	}

	for _, a := range jsArray {
		fmt.Println(a)
	}
	return goArray, nil
}
