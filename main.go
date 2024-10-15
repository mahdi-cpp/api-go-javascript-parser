package main

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_chat"
	"github.com/mahdi-cpp/api-go-javascript-parser/modelv2"
)

func main() {

	model.StartArrayParse()
	//repository.TestFunction()

	//model.StartParseViews()
	model.InitChat()
	model.InitPhotos()
	modelv2.InitPhotos()
	model_chat.InitModels()

	Run()
}
