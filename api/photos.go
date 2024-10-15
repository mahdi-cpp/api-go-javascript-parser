package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/modelv2"
)

func AddPhotosRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/gallery", func(context *gin.Context) {
		context.JSON(210, modelv2.RestPhotosV2())
	})

	route.GET("/all", func(context *gin.Context) {
		//model.StartParseViews()
		context.JSON(210, model.RestPhotos())
	})
}
