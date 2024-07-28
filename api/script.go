package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
)

func AddJavascriptRoutes(rg *gin.RouterGroup) {

	javascript := rg.Group("/script")

	javascript.GET("/test", func(context *gin.Context) {

		model.StartParseViews()
		//repository.TestFunction()

		context.JSON(210, model.RestAll())
	})

	javascript.GET("/feed", func(context *gin.Context) {

		model.StartParseViews()
		//repository.TestFunction()

		//context.JSON(210,
		context.JSON(210, model.RestAll())

	})

}
