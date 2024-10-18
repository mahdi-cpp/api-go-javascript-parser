package model_photos

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
)

type GalleryDTO struct {
	Photos []repository.PhotoBase `json:"photos"`
}

var galleryDTO GalleryDTO
var photoBaseArray []repository.PhotoBase

func GetGalleries(folder string) {

	var file = "data.txt"
	photos := repository.ReadOfFile(folder, file)
	var count = len(photos)
	var index = 0

	for i := 0; i < count; i++ {
		var photo = repository.PhotoBase{}
		photo = photos[index]
		photo.Key = -1
		galleryDTO.Photos = append(galleryDTO.Photos, photo)
		index++
	}
}
