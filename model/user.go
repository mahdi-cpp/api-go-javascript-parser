package model

import (
	"log"
	"os"
)

type Users struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var users []Users

func GetUsers() {
	users = []Users{}
	entries, err := os.ReadDir("/var/instagram/music/face/")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		users = append(users, Users{Name: "Name", Url: e.Name()})
	}
}
