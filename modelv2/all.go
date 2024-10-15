package modelv2

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-javascript-parser/repository"
	"github.com/mahdi-cpp/api-go-javascript-parser/utils"
	"math"
	"sync"
)

var root = "/"

func RestPhotosV2() map[string]any {
	return gin.H{
		"galleryDTO":          galleryDTO,
		"recentlyDTO":         recentlyDTO,
		"peopleDTO":           peopleDTO,
		"tripDTO":             tripDTO,
		"pinnedCollectionDTO": pinnedCollectionDTO,
		"albums":              albums,
		"shareAlbums":         shareAlbums,
		"cameraDTO":           cameraDTO,
	}
}

func InitPhotos() {

	GetGalleries("/var/cloud/bb/")
	GetRecently("/var/cloud/bb/")
	GetPeoples("/var/cloud/bb/")
	GetTrips("/var/cloud/bb/")
	GetPinned("/var/cloud/bb/")
	albums = GetAlbums("/var/cloud/bb/")
	shareAlbums = GetAlbums("/var/cloud/bb/")
	cameraDTO = GetCameras("/var/instagram/id/cut/")

	utils.GetCities()
	utils.GetNames()
}

type PhotoBaseCache struct {
	sync.RWMutex
	Cache map[int]repository.PhotoBase
}

func dp(value float32) float32 {
	if value == 0 {
		return 0
	}
	return float32(math.Ceil(float64(2.625 * value)))
}
