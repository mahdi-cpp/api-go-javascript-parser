package modelv2

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
)

type GalleryDTO struct {
	Photos      []repository.PhotoBase `json:"photos"`
	Titles      []string               `json:"titles"`
	SubTitles   []string               `json:"subTitles"`
	IconsUrl    []string               `json:"iconsUrl"`
	LargePhotos []string               `json:"largePhotos"`
}

var galleryDTO GalleryDTO
var photoBaseArray []repository.PhotoBase

func GetGalleries(folder string) {

	var file = "data.txt"
	photos := repository.ReadOfFile(folder, file)
	var count = len(photos)

	//if count > 50 {
	//	count = 50
	//}
	var index = 0

	for i := 0; i < count; i++ {
		var photo = repository.PhotoBase{}
		photo = photos[index]
		photo.Folder = folder
		photo.Key = -1
		galleryDTO.Photos = append(galleryDTO.Photos, photo)
		index++
	}

	galleryDTO.Titles = []string{"", "", "", "Videos", "Favourites", "Suggestions", "Crater Lake", "", "", "", "", ""}
	galleryDTO.SubTitles = []string{"", "MEDIA TYPES", "LIBRARY", "Mahdi", "TRIPS", "", "", "", "", ""}
	galleryDTO.IconsUrl = []string{"icons/camera_51.png", "icons/favourite_50.png", "icons/camera_photo_51.png", "icons/trip_50.png", "icons/trip_50.png"}
	galleryDTO.LargePhotos = []string{"",
		"",
		"",
		"id/cut/461979718_531724762937062_3561254644049376308_n.jpg",
		"id/girl/2019-05-14_15-36-41_UTC_2.jpg",
		"id/fa/8.jpg",
		"id/fa/sabihe__sb_1712847389_3343928730399948118_65847846068.jpg",
		"girl/sab/sabihe__sb_1715070340_3362576188340207639_65847846068 (1).jpg",
		"id/fa/farahmand_alipour_1623067994_2590804576664165813_201951656.jpg",
		"ik/PH183257.jpg"}
}
