package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
)

func StarScriptParse() {
	repository.ScriptParse()
}

func AddJavascriptRoutes(rg *gin.RouterGroup) {

	javascript := rg.Group("/javascript")

	javascript.GET("/images", func(context *gin.Context) {

		podcastDTO := repository.GetImages()

		context.JSON(210, gin.H{
			"javascriptDTO": podcastDTO,
		})
	})
}
