package themoviedb

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-movie/themoviedb/model"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func listMyMovies() {

	var root = "/media/mahdi/Movie"

	//files, err := os.ReadDir(root)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for _, file := range files {
	//	if file.IsDir() {
	//		directory(path, file)
	//	} else {
	//		//name, year := checkMovieName(file.Name())
	//		//if name != nil {
	//		//	//searchMovies(*name, *year)
	//		//}
	//	}
	//}

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {

			} else {
				name, year := checkMovieName(info.Name())
				if name != nil {
					searchMovies(*name, *year)
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func listMySeries() {

	series := []string{}

	var root = "/media/mahdi/My Passport/Series"
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}
			if info.IsDir() {

			} else {

				if strings.HasSuffix(info.Name(), ".mp4") || strings.HasSuffix(info.Name(), ".mkv") {
					//fmt.Println(info.Name())
				}

				name, _, _ := checkSeriesName(info.Name())
				if name != nil {
					found := slices.Contains(series, *name)
					if found == false {
						series = append(series, *name)
					}
					//fmt.Println(name, season, episode)
					//searchMovies(*name, *episode)
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("-------------------")
	for _, v := range series {
		//searchSeries(v)
		fmt.Println(v)
	}

	searchSeries("friends")
}

func listMyAnimations() {

	var root = "/media/mahdi/Photo/animation"

	i := 0

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {

			} else {
				name, year := checkMovieName(info.Name())
				if name != nil {
					i++
					fmt.Println(i, *name)
					fmt.Println("                                                                           Year:" + *year)
					searchAnimations(*name, *year)
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func searchMovies(name string, year string) {

	var tryCount = 4
	var finded = false
	var searchName = strings.ReplaceAll(name, " ", "%20")

searchAgain:
	tryCount++

	url := "https://api.themoviedb.org/3/search/movie?query=" + searchName + "&include_adult=false&language=en-US&page=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NDlkNWEwNGZhODNlMzkxZDcwNWY2NjY1NzdiM2RiYyIsInN1YiI6IjY2NmUxNThiZWE4MGFjNWViNTZiYTJhNSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ._kp_-lemNmcPiQzhKhb_UpV21AL5XOG7kkhO5qYC3-Q")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))

	var movies model.Movies
	err := json.Unmarshal([]byte(body), &movies)
	if err != nil {
		fmt.Println("searchMovies", "json.Unmarshal:", err)
		return
	}

	if len(movies.Results) > 0 {

		for _, result := range movies.Results {
			//fmt.Println("loop", result.Title, result.ReleaseDate)
			if strings.Contains(result.ReleaseDate, year) {
				fmt.Println(result.Id)
				fmt.Println(result.Title, result.ReleaseDate)
				fmt.Println(result.Overview)
				var url = "https://image.tmdb.org/t/p/original" + result.PosterPath
				fmt.Println(url)
				//fmt.Println(result.ReleaseDate, year)
				fmt.Println("------------------")
				//download(url, result.PosterPath)
				finded = true
				break
			}
		}

		fmt.Println("------------------------")

		//if strings.Compare(movies.Results[0].ReleaseDate, year) == 1 {
		//	fmt.Println(movies.Results[0].Title)
		//	var url = "https://image.tmdb.org/t/p/original" + movies.Results[0].PosterPath
		//	fmt.Println(url)
		//	fmt.Println("------------------")
		//} else {
		//	fmt.Println(movies.Results[0].ReleaseDate)
		//	fmt.Println(movies.Results[1].ReleaseDate)
		//	fmt.Println(year)
		//	fmt.Println("------------------")
		//}
	} else {
		fmt.Println("error", "searchMovies", name)
	}

	if finded == false && tryCount < 2 { //sometimes if keywords set short can find movies
		slice := strings.Split(name, " ")

		if len(slice) > 1 {
			fmt.Println("Try again ", name)

			slice = slice[:len(slice)-1]
			searchName = strings.Join(slice, "%20")
			fmt.Println(slice)

			var index = strings.LastIndex(name, " ")

			searchName = substr(searchName, 0, index)
			fmt.Println(searchName)

			goto searchAgain
		}
	}

}

func searchSeries(name string) {

	var searchName = strings.ReplaceAll(name, " ", "%20")
	url := "https://api.themoviedb.org/3/search/tv?query=" + searchName + "&include_adult=false&language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NDlkNWEwNGZhODNlMzkxZDcwNWY2NjY1NzdiM2RiYyIsInN1YiI6IjY2NmUxNThiZWE4MGFjNWViNTZiYTJhNSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ._kp_-lemNmcPiQzhKhb_UpV21AL5XOG7kkhO5qYC3-Q")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	var series model.Series
	err := json.Unmarshal([]byte(body), &series)
	if err != nil {
		fmt.Println("searchSeries", "json.Unmarshal:", err)
		return
	}

	if len(series.Results) > 0 {

		for _, result := range series.Results {
			//fmt.Println("loop", result.Title, result.ReleaseDate)
			//if strings.Contains(result.FirstAirDate, year) {
			fmt.Println(result.Id)
			//fmt.Println(result.Title, result.ReleaseDate)
			fmt.Println(result.Overview)
			var url = "https://image.tmdb.org/t/p/original" + result.PosterPath
			fmt.Println(url)
			//fmt.Println(result.ReleaseDate, year)
			fmt.Println("------------------")
			download(url, result.PosterPath)
			break
			//}
		}
		fmt.Println("------------------------")
	} else {
		fmt.Println("error", "searchMovies", name)
	}

}

func searchAnimations(name string, year string) {

	var tryCount = 4
	var finded = false
	var searchName = strings.ReplaceAll(name, " ", "%20")

searchAgain:
	tryCount++

	url := "https://api.themoviedb.org/3/search/movie?query=" + searchName + "&include_adult=false&language=en-US&page=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NDlkNWEwNGZhODNlMzkxZDcwNWY2NjY1NzdiM2RiYyIsInN1YiI6IjY2NmUxNThiZWE4MGFjNWViNTZiYTJhNSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ._kp_-lemNmcPiQzhKhb_UpV21AL5XOG7kkhO5qYC3-Q")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))

	var movies model.Movies
	err := json.Unmarshal([]byte(body), &movies)
	if err != nil {
		fmt.Println("searchMovies", "json.Unmarshal:", err)
		return
	}

	if len(movies.Results) > 0 {

		for _, result := range movies.Results {
			//fmt.Println("loop", result.Title, result.ReleaseDate)
			if strings.Contains(result.ReleaseDate, year) {
				fmt.Println(result.Id)
				fmt.Println(result.Title, result.ReleaseDate)
				fmt.Println(result.Overview)
				var url = "https://image.tmdb.org/t/p/original" + result.PosterPath
				fmt.Println(url)
				//fmt.Println(result.ReleaseDate, year)
				fmt.Println("------------------")
				download(url, result.PosterPath)
				finded = true
				break
			}
		}

		fmt.Println("------------------------")

	} else {
		fmt.Println("error", "searchMovies", name)
	}

	if finded == false && tryCount < 2 { //sometimes if keywords set short can find movies
		slice := strings.Split(name, " ")

		if len(slice) > 1 {
			fmt.Println("Try again ", name)

			slice = slice[:len(slice)-1]
			searchName = strings.Join(slice, "%20")
			fmt.Println(slice)

			var index = strings.LastIndex(name, " ")

			searchName = substr(searchName, 0, index)
			fmt.Println(searchName)

			goto searchAgain
		}
	}

}

func getMovieArtists(movie_id string) {

	url := "https://api.themoviedb.org/3/movie/" + movie_id + "/credits?language=en-US"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NDlkNWEwNGZhODNlMzkxZDcwNWY2NjY1NzdiM2RiYyIsInN1YiI6IjY2NmUxNThiZWE4MGFjNWViNTZiYTJhNSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ._kp_-lemNmcPiQzhKhb_UpV21AL5XOG7kkhO5qYC3-Q")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	var movieArtists model.MovieArtists
	err := json.Unmarshal([]byte(body), &movieArtists)
	if err != nil {
		fmt.Println("searchMovies", "json.Unmarshal:", err)
		return
	}

	if len(movieArtists.Cast) > 0 {
		for _, result := range movieArtists.Cast {

			if len(result.ProfilePath) > 5 {
				fmt.Println(result.Name)
				url := "https://image.tmdb.org/t/p/original" + result.ProfilePath
				fmt.Println(url)
				//download(url, result.ProfilePath)
			}

		}
	}

}

func searchPerson(name string) {

	name = strings.ReplaceAll(name, " ", "%20")

	url := "https://api.themoviedb.org/3/search/person?query=" + name + "&include_adult=false&language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NDlkNWEwNGZhODNlMzkxZDcwNWY2NjY1NzdiM2RiYyIsInN1YiI6IjY2NmUxNThiZWE4MGFjNWViNTZiYTJhNSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ._kp_-lemNmcPiQzhKhb_UpV21AL5XOG7kkhO5qYC3-Q")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

func searchExternal(movie_id string) {

	url := "https://api.themoviedb.org/3/movie/" + movie_id + "/external_ids"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NDlkNWEwNGZhODNlMzkxZDcwNWY2NjY1NzdiM2RiYyIsInN1YiI6IjY2NmUxNThiZWE4MGFjNWViNTZiYTJhNSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ._kp_-lemNmcPiQzhKhb_UpV21AL5XOG7kkhO5qYC3-Q")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
