package main

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/api"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_chat"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_drawer"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_parse"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_photos"
)

func main() {

	model_parse.StartArrayParse()
	//repository.TestFunction()

	//model_parse.StartParseViews()

	//model_parse.InitChat()
	//model_parse.InitPhotos()

	model_photos.InitPhotos()
	model_chat.InitModels()
	model_drawer.InitDrawers()

	api.ReadIcons()

	Run()
}
