package model

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func RestPhotos() map[string]any {
	return gin.H{
		"allGallery":       allGallery,
		"recentlyPhotos":   recentlyPhotos,
		"people":           people,
		"tripDTO":          tripDTO,
		"pinnedCollection": pinnedCollection,
		"albums":           galleryAlbums,
		"shareAlbums":      shareAlbums,
		"cameras":          galleryCameras,
	}
}

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	ProvinceId int    `json:"province_id"`
}

type Photo struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Orientation  int    `json:"orientation"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileType     string `json:"fileType"`
}

type Album struct {
	Name   string  `json:"name"`
	Photos []Photo `json:"photos"`
}

type People struct {
	Albums          []Album `json:"albums"`
	PhotoAnimations []Photo `json:"photoAnimations"`
}

type Trip struct {
	Name   string  `json:"name"`
	Date   string  `json:"date"`
	Photos []Photo `json:"photos"`
}
type TripDTO struct {
	Trips           []Trip  `json:"trips"`
	PhotoAnimations []Photo `json:"photoAnimations"`
}

var allGallery []Photo
var recentlyPhotos []Album

var pinnedCollection []Photo
var galleryAlbums []Album
var shareAlbums []Album
var galleryCameras []Album

var people People
var tripDTO TripDTO

var cities []City

func InitPhotos() {
	getCities()

	GetGalleries()
	GetRecently()
	GetPeoples()
	GetTrips()
	GetPinnedCollection()
	GetGalleryAlbums()
	GetShareAlbums()
	GetGalleryCameras()
}

func getCities() {
	var folder = "/data/"
	var file = "cities.json"
	cities = []City{}

	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cities); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//for _, city := range cities {
	//	fmt.Println(city.Name)
	//}
	//fmt.Println(len(cities))
}

func GetGalleries() {

	var folder = "/id/girl/"
	var file = "data.txt"
	allGallery = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&allGallery); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func GetRecently() {
	var folder = "/id/girl/"
	var file = "data.txt"
	var photos []Photo
	recentlyPhotos = []Album{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos) / 10
	var index = 0

	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = "Album " + strconv.Itoa(i)
		for j := 0; j < 10; j++ {
			album.Photos = append(album.Photos, photos[index])
			index++
		}
		recentlyPhotos = append(recentlyPhotos, album)
	}
}

func GetPeoples() {
	var folder = "/id/face/"
	var file = "data.txt"
	var photos []Photo
	people = People{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos) / 4
	var index = 0

	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = cities[i+50].Name
		for j := 0; j < 4; j++ {
			album.Photos = append(album.Photos, photos[index])
			index++
		}
		people.Albums = append(people.Albums, album)
	}

	GetPerson()
}

func GetPerson() {
	var folder = "/id/person/"
	var file = "data.txt"
	var photos []Photo

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos)
	var index = 0

	for i := 0; i < count; i++ {
		people.PhotoAnimations = append(people.PhotoAnimations, photos[index])
		index++
	}

}

func GetTrips() {

	var folder = "/id/trip/"
	var file = "data.txt"

	var photos []Photo
	tripDTO = TripDTO{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos)
	var index = 0

	for i := 0; i < count; i++ {
		var newTrip = Trip{}
		newTrip.Name = cities[i].Name
		newTrip.Date = "1403/05/14"
		for j := 0; j < 1; j++ {
			newTrip.Photos = append(newTrip.Photos, photos[index])
			index++
		}
		tripDTO.Trips = append(tripDTO.Trips, newTrip)
	}

	getTripAnimations()
}

func getTripAnimations() {
	var folder = "/id/trip2/"
	var file = "data.txt"
	var photos []Photo

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos)
	var index = 0

	for i := 0; i < count; i++ {
		tripDTO.PhotoAnimations = append(tripDTO.PhotoAnimations, photos[index])
		index++
	}
}

func GetPinnedCollection() {

	var folder = "/id/pinned/"
	var file = "data.txt"
	pinnedCollection = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&pinnedCollection); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func GetGalleryAlbums() {
	var folder = "/id/animal/"
	var file = "data.txt"
	var photos []Photo
	galleryAlbums = []Album{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos) / 10
	var index = 0

	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = "Album " + strconv.Itoa(i)
		for j := 0; j < 10; j++ {
			album.Photos = append(album.Photos, photos[index])
			index++
		}
		galleryAlbums = append(galleryAlbums, album)
	}
}

func GetShareAlbums() {
	var folder = "/id/go/"
	var file = "data.txt"
	var photos []Photo
	shareAlbums = []Album{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos) / 10
	var index = 0

	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = "Share Album " + strconv.Itoa(i)
		for j := 0; j < 10; j++ {
			album.Photos = append(album.Photos, photos[index])
			index++
		}
		shareAlbums = append(shareAlbums, album)
	}
}

func GetGalleryCameras() {
	var folder = "/id/tde/"
	var file = "data.txt"
	var photos []Photo
	galleryCameras = []Album{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var count = len(photos) / 10
	var index = 0

	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = "Album " + strconv.Itoa(i)
		for j := 0; j < 10; j++ {
			album.Photos = append(album.Photos, photos[index])
			index++
		}
		galleryCameras = append(galleryCameras, album)
	}
}
