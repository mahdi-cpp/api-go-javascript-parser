package model_chat

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
)

var storyDTO StoryDTO

type StoryDTO struct {
	Caption string               `json:"name"`
	Avatar  repository.PhotoBase `json:"avatar"`
	Stories []Story              `json:"stories"`
}

type Story struct {
	Name  string               `json:"name"`
	Photo repository.PhotoBase `json:"photo"`
}

func GetStory(folder string, avatar string) StoryDTO {

	var file = "data.txt"
	var photos = repository.ReadOfFile(folder, file)
	var count = len(photos)
	var storyDTO StoryDTO

	if count > 50 {
		count = 50
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
	storyDTO.Avatar = photo

	for i := 0; i < count; i++ {
		var story Story
		if nameIndex >= len(utils.FackNames) {
			nameIndex = 0
		}

		story.Name = utils.FackNames[nameIndex]

		story.Photo = photos[index]
		story.Photo.Key = -1
		story.Photo.ThumbSize = 270
		story.Photo.Circle = true
		story.Photo.PaintWidth = dp(70)
		story.Photo.PaintHeight = dp(70)

		storyDTO.Stories = append(storyDTO.Stories, story)
		nameIndex++
		index++
	}

	index = 0

	return storyDTO
}
