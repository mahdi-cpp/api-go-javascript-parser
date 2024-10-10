package model

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
)

var root = "/var/instagram"

func RestChats() map[string]any {
	return gin.H{
		"camera":         camera,
		"movies":         movies,
		"series":         series,
		"stories":        stories,
		"albums":         galleryAlbums,
		"instagramPost1": instagramPost1,
		"instagramPost2": instagramPost2,
		"instagramPost3": instagramPost3,
		"recentlyDTOS":   recentlyDTOS,
		"cameras":        galleryCameras,
		"chatDaysDTO":    chatDaysDTO,
		"pagerPhotos":    pagerPhotos,
		"questionVoices": questionVoices,
		"pdfs":           pdfs,
		"maps":           maps,
		"notes":          notes,
		"musics":         musics,
		"posts":          posts,
		"avatars":        avatars,

		//"functions":     RestFunctions(),
		//"scrollViews":   scrollViews,
		//"textBoxes":     textBoxes,
		//"textViews":     textViews,
		//"images":        images,
		//"buttons":       button,
		//"switchButtons": switchButtons,
		//"sliderViews":   sliderViews,
		//"chartViews":    chartViews,
		//"textMessages":  textMessages,
		//"textInputs":    textInputs,
	}
}

type Chat struct {
	ChatId         int    `json:"chatId"`
	ChatName       string `json:"chatName"`
	AvatarUrl      string `json:"avatarUrl"`
	UnreadMessages int    `json:"unreadMessages"`
	LastMessage    string `json:"lastMessage"`
	LastUpdate     int    `json:"lastUpdate"`
}

var chats []Chat

var movies []Photo
var series []Photo

var stories []Photo
var chatDaysDTO []ChatDays

var avatars []Photo
var posts []Photo
var musics []Photo
var notes []Photo
var albums []Photo
var maps []Photo
var pdfs []Photo
var questionVoices []Photo
var pagerPhotos []Photo
var instagramPhotos []Photo

var instagramPost1 []Photo
var instagramPost2 []Photo
var instagramPost3 []Photo
var camera []Photo

type ChatDays struct {
	Name   string  `json:"name"`
	Photos []Photo `json:"photos"`
}

func InitChat() {
	GetStories()
	GetMovies()
	GetSeries()
	GetPdfs()

	GetAvatars()
	GetPosts()
	GetMusics()
	GetNotes()
	GetAlbums()
	GetMaps()
	GetQuestionVoices()
	GetTestPager()
	GetDaysAlbums()
	ChatInit()

	GetInstagramPost1()
	GetInstagramPost2()
	GetInstagramPost3()

	GetCamera()
}

func ChatInit() {
	GetChats()
}

func GetChats() {
	chats = []Chat{}
	var path = "/face2/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		chats = append(chats, Chat{ChatId: 1, AvatarUrl: path + e.Name()})
	}
}

func GetInstagramPost1() {
	var folder = "/id/girl/"
	var file = "data.txt"
	instagramPost1 = []Photo{}

	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&instagramPost1); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func GetInstagramPost2() {
	var folder = "/id/trip3/"
	var file = "data.txt"
	instagramPost2 = []Photo{}

	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&instagramPost2); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func GetInstagramPost3() {
	var folder = "/id/go/"
	var file = "data.txt"
	instagramPost3 = []Photo{}

	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&instagramPost3); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func GetCamera() {
	var folder = "/id/go/"
	var file = "data.txt"
	camera = []Photo{}

	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&camera); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func GetStories() {
	var folder = "/movie/casts/"
	var file = "data.txt"
	stories = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&stories); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func GetMovies() {
	var folder = "/movie/movie/"
	var file = "data.txt"
	movies = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&movies); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func GetSeries() {
	var folder = "/movie/series/"
	var file = "data.txt"
	series = []Photo{}

	// Open the file for reading
	f, err := os.Open(root + folder + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Create a JSON decoder and decode the data into the slice
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&series); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func GetPdfs() {
	pdfs = []Photo{}
	var path = "/pdf/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		pdfs = append(pdfs, Photo{Name: "Name", Url: path + e.Name()})
	}
}

func GetDaysAlbums() {
	var folder = "/id/girl/"
	var file = "data.txt"
	var photos []Photo
	chatDaysDTO = []ChatDays{}

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

	var count = len(photos) / 5
	var index = 0

	for i := 0; i < count; i++ {
		var album = ChatDays{}
		album.Name = "Album " + strconv.Itoa(i)
		for j := 0; j < 5; j++ {
			album.Photos = append(album.Photos, photos[index])
			index++
		}
		chatDaysDTO = append(chatDaysDTO, album)
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
	var path = "/face2/"
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

func GetMaps() {
	maps = []Photo{}
	var path = "/map/"
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
