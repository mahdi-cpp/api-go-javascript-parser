package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/model"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_drawer"
	"github.com/mahdi-cpp/api-go-javascript-parser/model_photos"
)

func AddPhotosRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/gallery", func(context *gin.Context) {
		context.JSON(210, model_photos.RestPhotosV2())
	})

	route.GET("/all", func(context *gin.Context) {
		context.JSON(210, model.RestPhotos())
	})

	route.GET("/people", func(context *gin.Context) {
		context.JSON(210, model_drawer.RestPeople())
	})

	route.GET("/music", func(context *gin.Context) {
		context.JSON(210, model_drawer.RestMusic())
	})
}
