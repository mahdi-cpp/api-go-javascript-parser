package model

import (
	"fmt"
	"github.com/disintegration/imaging"
	"log"
	"os"
)

var root = "/var/instagram"

type Users struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

var users []Users
var avatars []Users
var posts []Users
var musics []Users
var movies []Users
var gallery []Users
var albums []Users
var maps []Users
var pdfs []Users
var questionVoices []Users

func GetQuestionVoices() {
	questionVoices = []Users{}
	var path = "/student/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		questionVoices = append(questionVoices, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetPdfs() {
	pdfs = []Users{}
	var path = "/pdf/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		pdfs = append(pdfs, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetMaps() {
	maps = []Users{}
	var path = "/map/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		maps = append(maps, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetAlbums() {
	albums = []Users{}
	var path = "/nature/01/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetGallery() {
	gallery = []Users{}
	var path = "/ios/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		gallery = append(gallery, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetMovies() {
	movies = []Users{}
	var path = "/posters/thumbnail_2/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		movies = append(movies, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetMusics() {
	musics = []Users{}
	var path = "/camera/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		musics = append(musics, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetAvatars() {
	avatars = []Users{}
	var path = "/person2/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		avatars = append(avatars, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetUsers() {
	users = []Users{}
	var path = "/face/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		//width, height := getImageDimension(root + path + e.Name())
		users = append(users, Users{Name: "Name", Url: path + e.Name()})
	}
}

func GetPosts() {
	posts = []Users{}
	var path = "/behance/03/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		width, height := getImageDimension(root + path + e.Name())
		posts = append(posts, Users{Name: "Name", Url: path + e.Name(), Width: width, Height: height})
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
