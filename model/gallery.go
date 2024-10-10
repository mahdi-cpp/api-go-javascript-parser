package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

var id = 0

var ThumbsUrl = []string{
	"_70.jpg",
	"_135.jpg",
	"_270.jpg",
	"_540.jpg",
}

func RestPhotos() map[string]any {
	return gin.H{
		"galleryDTO":        galleryDTO,
		"recentlyDTOS":      recentlyDTOS,
		"peopleDTO":         peopleDTO,
		"tripDTO":           tripDTO,
		"pinnedCollections": pinnedCollections,
		"albums":            galleryAlbums,
		"shareAlbums":       shareAlbums,
		"cameras":           galleryCameras,
	}
}

type Photo struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Url          string   `json:"url"`
	ThumbnailUrl string   `json:"thumbnailUrl"`
	Orientation  int      `json:"orientation"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	FileType     string   `json:"fileType"`
	ThumbsUrl    []string `json:"thumbsUrl"`
}

type GalleryDTO struct {
	Photos      []Photo  `json:"photos"`
	Titles      []string `json:"titles"`
	SubTitles   []string `json:"subTitles"`
	IconsUrl    []string `json:"iconsUrl"`
	LargePhotos []string `json:"largePhotos"`
}

type Recently struct {
	Name   string  `json:"name"`
	Photos []Photo `json:"photos"`
}

type Album struct {
	Name   string  `json:"name"`
	Name2  string  `json:"name2"`
	Photos []Photo `json:"photos"`
}

type PersonGroup struct {
	Names  []string `json:"names"`
	Photos []Photo  `json:"photos"`
}
type PeopleDTO struct {
	PersonGroups    []PersonGroup `json:"personGroups"`
	PhotoAnimations []Photo       `json:"photoAnimations"`
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

type PinnedCollection struct {
	Name  string `json:"name"`
	Photo Photo  `json:"photo"`
}

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	ProvinceId int    `json:"province_id"`
}

var galleryDTO GalleryDTO
var recentlyDTOS []Recently

var pinnedCollections []PinnedCollection
var galleryAlbums []Album
var shareAlbums []Album
var galleryCameras []Album

var peopleDTO PeopleDTO
var tripDTO TripDTO

var cities []City
var names []string

func InitPhotos() {

	GetGalleries()
	GetRecently()
	GetPeoples()
	GetTrips()
	GetPinnedCollection()
	GetGalleryAlbums()
	GetShareAlbums()
	GetGalleryCameras()

	getCities()
	getNames()
}

func GetGalleries() {

	var folder = "/id/bb/"
	var file = "data.txt"
	galleryDTO.Photos = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&galleryDTO.Photos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
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

func GetRecently() {
	var folder = "/id/girl/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo
	recentlyDTOS = []Recently{}

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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2) / 5
	var index = 0
	for i := 0; i < count; i++ {
		var album = Recently{}
		album.Name = "Album " + strconv.Itoa(i)
		for j := 0; j < 5; j++ {
			album.Photos = append(album.Photos, photos2[index])
			index++
		}
		recentlyDTOS = append(recentlyDTOS, album)
	}
}

func GetPeoples() {
	var folder = "/id/girl/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo
	peopleDTO = PeopleDTO{}

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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2) / 4
	var index = 0
	var nameIndex = 0
	for i := 0; i < count; i++ {
		var personGroup = PersonGroup{}

		if nameIndex >= len(FackNames) {
			nameIndex = 0
		}

		for j := 0; j < 4; j++ {
			if nameIndex >= len(FackNames) {
				nameIndex = 0
			}
			personGroup.Names = append(personGroup.Names, FackNames[nameIndex])
			personGroup.Photos = append(personGroup.Photos, photos2[index])
			index++
			nameIndex++
		}
		peopleDTO.PersonGroups = append(peopleDTO.PersonGroups, personGroup)
	}

	GetPerson()
}

func GetPerson() {
	var folder = "/id/bb/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo

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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2)
	var index = 0
	for i := 0; i < count; i++ {
		peopleDTO.PhotoAnimations = append(peopleDTO.PhotoAnimations, photos2[index])
		index++
	}

}

func GetTrips() {

	var folder = "/id/bb/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo
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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos)
	var index = 0
	var nameIndex = 0
	for i := 0; i < count; i++ {
		var newTrip = Trip{}
		if nameIndex >= len(FackTrips) {
			nameIndex = 0
		}
		newTrip.Name = FackTrips[nameIndex]
		newTrip.Date = "1403/05/14"
		for j := 0; j < 1; j++ {
			newTrip.Photos = append(newTrip.Photos, photos2[index])
			index++
		}
		tripDTO.Trips = append(tripDTO.Trips, newTrip)
		nameIndex++
	}

	getTripAnimations()
}

func getTripAnimations() {
	var folder = "/id/bb/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo

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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2)
	var index = 0
	for i := 0; i < count; i++ {
		tripDTO.PhotoAnimations = append(tripDTO.PhotoAnimations, photos2[index])
		index++
	}
}

func GetPinnedCollection() {

	var folder = "/id/pinned/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo

	pinnedCollections = []PinnedCollection{}
	a := []string{"Map", "Electronic", "Tiny House", "Projects", "Sensors", "Mechanic", "History", "Fave Clip", "Recently Saved",
		"Screen Records", "Office projects", "Map", "Recently Saved", "Electronic", "FaveClip"}

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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var index = 0
	for i := 0; i < len(photos2); i++ {
		var collection = PinnedCollection{}
		collection.Name = a[index]
		collection.Photo = photos2[index]
		pinnedCollections = append(pinnedCollections, collection)
		index++
	}
}

func GetGalleryAlbums() {
	var folder = "/id/bb/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo
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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2) / 10
	var index = 0
	var nameIndex = 0
	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = GalleryAlbumTitles[nameIndex]
		album.Name2 = GalleryAlbumTitles[nameIndex+1]
		for j := 0; j < 10; j++ {
			album.Photos = append(album.Photos, photos2[index])
			index++
		}
		galleryAlbums = append(galleryAlbums, album)
		nameIndex += 2
	}
}

func GetShareAlbums() {
	var folder = "/id/go/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo
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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2) / 10
	var index = 0
	var nameIndex = 0
	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = ShareAlbumTitles[nameIndex]
		album.Name2 = ShareAlbumTitles[nameIndex+1]
		for j := 0; j < 10; j++ {
			album.Photos = append(album.Photos, photos2[index])
			index++
		}
		shareAlbums = append(shareAlbums, album)
		nameIndex += 2
	}
}

func GetGalleryCameras() {
	var folder = "/id/ali/"
	var file = "data.txt"
	var photos []Photo
	var photos2 []Photo
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

	for _, photo := range photos {
		for i := 0; i < len(ThumbsUrl); i++ {
			photo.ThumbsUrl = append(photo.ThumbsUrl, photo.ThumbnailUrl+photo.Name+ThumbsUrl[i])
		}
		photos2 = append(photos2, photo)
	}

	var count = len(photos2) / 5
	var index = 0

	for i := 0; i < count; i++ {
		var album = Album{}
		album.Name = GalleryCameraTitles[i]
		for j := 0; j < 5; j++ {
			album.Photos = append(album.Photos, photos2[index])
			index++
		}
		galleryCameras = append(galleryCameras, album)
	}
}

func getNames() {

	filename := root + "/data/name.txt"

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after reading

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line and append it to the slice
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
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
}
