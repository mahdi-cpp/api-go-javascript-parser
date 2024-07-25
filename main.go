package main

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
)

func main() {
	//repository.StartScriptParse()
	repository.TestFunction()
	Run()

	//vm := goja.New()
	//data, err := os.ReadFile("web/test5.js")
	//if err != nil {
	//	panic(err)
	//}
	//src := string(data)
	//_, err = vm.RunString(src)
	//if err != nil {
	//	panic(err)
	//}

	//result, err := vm.RunString("new Music(\"Ali\").show();")
	//
	//if err != nil {
	//	fmt.Println("Error calling JavaScript function:", err)
	//	return
	//}
	//fmt.Println("Result of calling function: ", result.Export())

}
