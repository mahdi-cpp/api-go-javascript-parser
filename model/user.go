package model

import (
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"log"
	"os"
	"strings"
)

var root = "/var/instagram"

var users []Photo
var avatars []Photo
var posts []Photo
var musics []Photo
var movies []Photo
var notes []Photo
var albums []Photo
var maps []Photo
var pdfs []Photo
var questionVoices []Photo
var pagerPhotos []Photo
var instagramPhotos []Photo
var gallery []Photo

func InitChat() {
	GetGallery()
	GetUsers()
	GetAvatars()
	GetPosts()
	GetMusics()
	GetMovies()
	GetNotes()
	GetAlbums()
	GetMaps()
	GetPdfs()
	GetQuestionVoices()
	GetTestPager()
	GetInstagram()
	ChatInit()
}

func GetUsers() {
	//users = []Photo{}
	//var path = "/id/6/"
	//var smallPath = "/id/6/small/"
	//entries, err := os.ReadDir(root + path)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, e := range entries {
	//	if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
	//		//width, height := getImageDimension(root + path + e.Name())
	//		users = append(users, Photo{Name: "Name", TinyUrl: smallPath + e.Name(), Url: path + e.Name()})
	//	}
	//}

	var folder = "/id/girl/"
	var file = "data.txt"
	users = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&users); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//for _, ali := range users {
	//	fmt.Println(ali.FileType)
	//}
}

func GetGallery() {
	gallery = []Photo{}
	var path = "/face/small/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			gallery = append(gallery, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetInstagram() {
	instagramPhotos = []Photo{}
	var path = "/map/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			instagramPhotos = append(instagramPhotos, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetTestPager() {
	pagerPhotos = []Photo{}
	var path = "/id/fa/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			pagerPhotos = append(pagerPhotos, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetQuestionVoices() {
	questionVoices = []Photo{}
	var path = "/id/ali/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			questionVoices = append(questionVoices, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetPdfs() {
	pdfs = []Photo{}
	var path = "/id/ali/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		pdfs = append(pdfs, Photo{Name: "Name", Url: path + e.Name()})
	}
}

func GetMaps() {
	maps = []Photo{}
	var path = "/id/me/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			maps = append(maps, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetAlbums() {
	albums = []Photo{}
	var path = "/snap/gallery/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			albums = append(albums, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetNotes() {
	notes = []Photo{}
	var path = "/ios/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			notes = append(notes, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetMovies() {
	movies = []Photo{}
	var path = "/posters/movies-select/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			movies = append(movies, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetMusics() {
	musics = []Photo{}
	var path = "/music/albums/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			musics = append(musics, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetAvatars() {
	avatars = []Photo{}
	var path = "/person2/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			avatars = append(avatars, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func GetPosts() {
	posts = []Photo{}
	var path = "/tinyhome/05/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".jpg") || strings.HasSuffix(e.Name(), ".jpeg") {
			//width, height := getImageDimension(root + path + e.Name())
			posts = append(posts, Photo{Name: "Name", Url: path + e.Name()})
		}
	}
}

func getImageDimension(imagePath string) (int, int) {
	img, err := imaging.Open(imagePath) // Replace "image.jpg" with the path to your image file
	if err != nil {
		fmt.Println("Error opening image:", err)
		return 0, 0
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	fmt.Printf("Image width: %d\n", width)
	fmt.Printf("Image height: %d\n", height)
	return width, height
}
