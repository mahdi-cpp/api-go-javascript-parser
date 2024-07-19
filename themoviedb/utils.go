package themoviedb

import (
	"bufio"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var numberOnly = regexp.MustCompile("^[0-9-_!@#$%^&*()]*$")

func isOnlyNumber(str string) bool {
	return numberOnly.MatchString(str)
}

func checkMovieName(fileName string) (*string, *string) {
	debug := false
	var funcName = "checkName:"

	var name = strings.ToLower(fileName)
	var year = 1900

	if strings.HasSuffix(name, ".mp4") || strings.HasSuffix(name, ".mkv") {

	} else {
		if debug == true {
			fmt.Println(funcName, "filetype must be .mp4 or .mkv suffix -->", "{"+name+"}")
		}
		return nil, nil
	}

	name = strings.ReplaceAll(name, ".", " ")
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "(", "")
	name = strings.ReplaceAll(name, ")", "")
	name = strings.ReplaceAll(name, "!", "")
	name = strings.ReplaceAll(name, "@", "")
	name = strings.ReplaceAll(name, "#", "")
	name = strings.ReplaceAll(name, "%", "")
	name = strings.ReplaceAll(name, "&", "")

	//fmt.Println(name)

	if isOnlyNumber(fileName) {
		if debug == true {
			fmt.Println(funcName, "filetype must be contain alphabetic characters -->", "{"+name+"}")
		}
		return nil, nil
	}

	for i := 1900; i < 2050; i++ {
		if strings.Contains(name, strconv.Itoa(i)) {
			year = i
			index := strings.Index(name, strconv.Itoa(year))
			if index > 2 {
				name = substr(name, 0, index)
				var year2 = strconv.Itoa(year)
				return &name, &year2
			} else {
				if debug == true {
					fmt.Println(funcName, "filename name and year not valid -->", "{"+name+"}")
				}
				return nil, nil
			}
		}
	}

	//fmt.Println(funcName, "FileName is Valid -->", "{"+name+"}")

	return nil, nil
}

func checkSeriesName(fileName string) (*string, *string, *string) {

	var funcName = "checkName:"

	var name = strings.ToLower(fileName)
	var season = "01"
	var episode = "01"

	if strings.HasSuffix(name, ".mp4") || strings.HasSuffix(name, ".mkv") {
		name = strings.ReplaceAll(name, ".", " ")
		name = strings.ReplaceAll(name, "_", " ")
		name = strings.ReplaceAll(name, "(", "")
		name = strings.ReplaceAll(name, ")", "")
		name = strings.ReplaceAll(name, "!", "")
		name = strings.ReplaceAll(name, "@", "")
		name = strings.ReplaceAll(name, "#", "")
		name = strings.ReplaceAll(name, "%", "")
		name = strings.ReplaceAll(name, "&", "")
	} else {
		//fmt.Println(funcName, "filetype must be .mp4 or .mkv suffix -->", "{"+name+"}")
		return nil, nil, nil
	}

	//fmt.Println(name)

	if isOnlyNumber(fileName) {
		//fmt.Println(funcName, "filetype must be contain alphabetic characters -->", "{"+name+"}")
		return nil, nil, nil
	}

	for i := 1; i < 30; i++ {
		var ss = "s0" + strconv.Itoa(i)
		if strings.Contains(name, ss) {
			index := strings.Index(name, ss)
			episode = substr(name, index+3, 3)
			if index > 5 {
				name = substr(name, 0, index)
				//fmt.Println(name, ss, episode)
			} else {
				fmt.Println(funcName, "filename name and season and episode not valid -->", "{"+name+"}")
				return nil, nil, nil
			}
			return &name, &season, &episode
		}
	}

	return nil, nil, nil
}

func download(url string, name string) {

	//if isExistFile(name) {
	//	return
	//}

	name = strings.ReplaceAll(name, "/", "")

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("/var/instagram/posters/artists/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}

func isExistFile(fileName string) bool {
	files, err := ioutil.ReadDir("/var/instagram/posters/animation")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		if strings.Contains(fileName, file.Name()) {
			fmt.Println("File is exist:", file.Name())
			return true
		}
	}
	return false
}

func readOfFile() {
	file, err := os.Open("/home/mahdi/photos/ali.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var url = "https://image.tmdb.org/t/p/original" + scanner.Text()
		fmt.Println(url)
		//fmt.Println("....")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
//
//	gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func directory(path string, directory fs.DirEntry) {

	var root = path + "/" + directory.Name()
	fmt.Println("is Directory", root)
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func createThump(photo string) {

	// input files
	path := "/var/instagram/01/"
	file := path + photo

	fi, err := os.Stat(file)
	if err != nil {
		return
	}

	var imageSizeKB int = 0
	imageSizeKB = int(fi.Size() / 1024)

	var width = 0
	var height = 0

	var w = 200
	var h = 300

	if reader, err := os.Open(filepath.Join(path, photo)); err == nil {
		defer reader.Close()
		im, _, err := image.DecodeConfig(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", photo, err)
			return
		}
		fmt.Println("%s %d %d\n", photo, im.Width, im.Height)

		width = im.Width
		height = im.Height

		if imageSizeKB >= 1000 {
			for i := 20; i >= 9; i-- {
				if imageSizeKB > (100 * i) {
					w = width / (i * 1)
					h = height / (i * 1)
					break
				}
			}
		} else {
			for i := 2000; i > 100; i -= 100 {
				if width > i {
					w = width / (i / 200)
					h = height / (i / 200)
					break
				}
			}
		}

	} else {
		fmt.Println("Impossible to open the file:", err)
	}

	if imageSizeKB < 20 {
		copyImage(path+photo, path+"/thumbnail_100/"+photo)
	} else {
		// load images and make 100x100 thumbnails of them
		var thumbnail image.Image
		img, err := imaging.Open(file)
		if err != nil {
			panic(err)
		}
		thumbnail = imaging.Thumbnail(img, w, h, imaging.CatmullRom)

		// create a new blank image
		dst := imaging.New(w, h, color.NRGBA{0, 0, 0, 0})

		// paste thumbnails into the new image side by side
		dst = imaging.Paste(dst, thumbnail, image.Pt(0, 0))

		// save the combined image to file
		err = imaging.Save(dst, path+"/thumbnail_100/"+photo)
		if err != nil {
			panic(err)
		}
	}

}

func copyImage(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
