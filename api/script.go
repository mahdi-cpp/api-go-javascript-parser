package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
)

func AddJavascriptRoutes(rg *gin.RouterGroup) {

	javascript := rg.Group("/script")

	javascript.GET("/feed", func(context *gin.Context) {

		repository.StartScriptParse()
		//repository.TestFunction()

		context.JSON(210,
			gin.H{
				"functions":     repository.RestFunctions(),
				"textBoxes":     repository.RestTextBoxes(),
				"textViews":     repository.RestTextViews(),
				"images":        repository.RestImages(),
				"circleButtons": repository.RestCircleButtons(),
				"switchButtons": repository.RestSwitchButtons(),
				"sliderViews":   repository.RestSliderView(),
				"chartViews":    repository.RestChartViews(),
			})
	})

}
