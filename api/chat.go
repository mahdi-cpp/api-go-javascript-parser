package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_chat"
)

func AddJavascriptRoutes(rg *gin.RouterGroup) {

	javascript := rg.Group("/chat")

	//javascript.GET("/chat", func(context *gin.Context) {
	//	//model_parse.StartParseViews()
	//	context.JSON(210, model_parse.RestChats())
	//})
	//
	//javascript.GET("/chats", func(context *gin.Context) {
	//	//model_parse.StartParseViews()
	//	context.JSON(210, model_parse.RestChats())
	//})

	javascript.GET("/chatV2", func(context *gin.Context) {
		context.JSON(210, model_chat.RestChatV2())
	})

}
