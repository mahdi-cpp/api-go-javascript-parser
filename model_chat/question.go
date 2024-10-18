package model_chat

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
)

var questionVoiceDTO QuestionVoiceDTO

type QuestionVoiceDTO struct {
	Caption        string          `json:"name"`
	QuestionVoices []QuestionVoice `json:"questionVoices"`
}

type QuestionVoice struct {
	Name  string               `json:"name"`
	Photo repository.PhotoBase `json:"photo"`
}

func GetQuestionVoices(folder string) QuestionVoiceDTO {

	var file = "data.txt"
	var photos = repository.ReadOfFile(folder, file)
	var count = len(photos)
	var dto QuestionVoiceDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var voice QuestionVoice
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		voice.Name = utils.MovieNames[nameIndex]

		voice.Photo = photos[index]
		voice.Photo.Key = -1
		voice.Photo.ThumbSize = 270
		voice.Photo.Circle = true
		voice.Photo.PaintWidth = dp(88)
		voice.Photo.PaintHeight = dp(88)

		dto.QuestionVoices = append(dto.QuestionVoices, voice)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
