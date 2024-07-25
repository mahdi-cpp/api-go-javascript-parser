package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
)

func AddJavascriptRoutes(rg *gin.RouterGroup) {

	javascript := rg.Group("/script")

	javascript.GET("/feed", func(context *gin.Context) {

		repository.StartScriptParse()
		repository.TestFunction()

		context.JSON(210,
			gin.H{
				"textViews":     repository.GetTextViews(),
				"images":        repository.GetImages(),
				"circleButtons": repository.GetCircleButtons(),
				"switchButtons": repository.GetSwitchButtons(),
				"sliderViews":   repository.GetSliderView(),
				"functions":     repository.RestFunctions(),
			})
	})

}
