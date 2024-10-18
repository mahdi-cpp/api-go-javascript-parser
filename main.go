package main

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/api"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_chat"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_drawer"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_photos"
)

func main() {

	model.StartArrayParse()
	//repository.TestFunction()

	//model.StartParseViews()

	//model.InitChat()
	//model.InitPhotos()

	model_photos.InitPhotos()
	model_chat.InitModels()
	model_drawer.InitDrawers()

	api.ReadIcons()

	Run()
}
