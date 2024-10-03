package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
)

func AddJavascriptRoutes(rg *gin.RouterGroup) {

	javascript := rg.Group("/chat")

	javascript.GET("/chat", func(context *gin.Context) {
		model.StartParseViews()
		context.JSON(210, model.RestChat())
	})

	javascript.GET("/chats", func(context *gin.Context) {
		model.StartParseViews()
		context.JSON(210, model.RestChats())
	})

}
