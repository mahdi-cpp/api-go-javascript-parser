package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
)

func AddPhotosRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/all", func(context *gin.Context) {
		model.StartParseViews()
		context.JSON(210, model.RestPhotos())
	})
}
