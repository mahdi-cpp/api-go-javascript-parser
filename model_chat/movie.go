package model_chat

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
)

var movieDTO MovieDTO

type MovieDTO struct {
	Caption string  `json:"name"`
	Movies  []Movie `json:"movies"`
}

type Movie struct {
	Name  string               `json:"name"`
	Photo repository.PhotoBase `json:"photo"`
}

func GetMovies(folder string) MovieDTO {

	var file = "data.txt"
	var photos = repository.ReadOfFile(folder, file)
	var count = len(photos)
	var dto MovieDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var movie Movie
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		movie.Name = utils.MovieNames[nameIndex]

		movie.Photo = photos[index]
		movie.Photo.Key = -1
		movie.Photo.ThumbSize = 270
		movie.Photo.PaintWidth = dp(70)
		movie.Photo.PaintHeight = dp(120)

		dto.Movies = append(dto.Movies, movie)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
