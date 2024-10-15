package model_chat

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
)

type InstagramPostDTO struct {
	Caption         string                 `json:"name"`
	Avatar          repository.PhotoBase   `json:"avatar"`
	InstagramPhotos []repository.PhotoBase `json:"instagramPhotos"`
}

var instagramPostDTO1 InstagramPostDTO
var instagramPostDTO2 InstagramPostDTO
var instagramPostDTO3 InstagramPostDTO

func GetInstagram(folder string, avatar string) InstagramPostDTO {

	var file = "data.txt"
	var photos = repository.ReadOfFile(folder, file)
	var count = len(photos)
	var localInstagramPostDTO InstagramPostDTO

	if count > 10 {
		count = 10
	}

	var index = 0
	var nameIndex = 0

	var photo = repository.PhotoBase{}
	photo.Name = avatar
	photo.FileType = ".jpg"
	photo.Width = 200
	photo.Height = 200
	photo.ThumbSize = 270
	photo.Crop = true
	photo.Round = int(dp(20))
	photo.Key = -1
	photo.PaintWidth = dp(100)
	photo.PaintHeight = dp(100)
	localInstagramPostDTO.Avatar = photo

	for i := 0; i < count; i++ {
		if nameIndex >= len(utils.FackNames) {
			nameIndex = 0
		}

		var photo = repository.PhotoBase{}
		photo = photos[index]
		photo.ThumbSize = 270
		photo.Crop = true
		photo.Round = int(dp(10))
		photo.Key = -1
		//photo.PaintWidth = photoSize
		//photo.PaintHeight = photoSize

		localInstagramPostDTO.InstagramPhotos = append(localInstagramPostDTO.InstagramPhotos, photo)

		nameIndex++
		index++
	}

	index = 0

	return localInstagramPostDTO
}
